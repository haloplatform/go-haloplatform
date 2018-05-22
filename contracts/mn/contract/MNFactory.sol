pragma solidity ^0.4.18;

import "./MNInstance.sol";

/** @title MN proposed market SC */
contract MNFactory {
  /// Set minimum HALO of a share
  uint private constant MIN_HALO_SHARE_K = 1 ether;

  /// Node state
  uint8 private constant MN_STATE_INVALID_K       = 0;
  uint8 private constant MN_STATE_INITIALIZING_K  = 1;
  uint8 private constant MN_STATE_DEPOSITED_K     = 2;

  /// Event type
  uint8 private constant MN_EVT_InitializeMN      = 1;
  uint8 private constant MN_EVT_AdminCancel       = 2;
  uint8 private constant MN_EVT_AdminAccept       = 3;
  uint8 private constant MN_EVT_FinalizeMN        = 4;
  uint8 private constant MN_EVT_AutoCancel        = 5;

  struct TierInfo { // Struct
    uint    COST;
    uint16  LIMIT;
    uint16  count;
    uint16  activatedCount;
    uint    PoW;
  }

  struct NodeInfo {
    uint32    id;       /// Human-readable ID
    uint8     tier;             /// Tier type
    uint8     state;
    bool      is_private;
    bool      is_halo_hosting;
  }

  struct SelectedNodeList {
    address[]   tier1;
    address[]   tier2;
    address[]   tier3;
    address[]   tier4;
  }

  address     private  m_addrAdmin;
  address     private  m_mnAdminAddr;
  address     private  m_mnMarketplaceAddr;
  address     private  m_mnRewardAddr;
  
  uint32      private  m_ui_idGen;
  TierInfo[4] private   m_tierSpec;
  address[]   private   m_nodeAddrList;
//  mapping (address => NodeInfo) private  m_nodeMap;

  event Event_MNFactory(address indexed _nodeAddr, uint32 _nodeId, uint8 _eventType);
  
  /**
   * Constructor
   */
  function MNFactory() public {
    // constructor
    m_addrAdmin = msg.sender;
    /// Tier cost
    m_tierSpec[0].COST = 500 ether;
    m_tierSpec[1].COST = 1000 ether;
    m_tierSpec[2].COST = 2500 ether;
    m_tierSpec[3].COST = 7500 ether;
    /// Limitation
    m_tierSpec[0].LIMIT = 5000;
    m_tierSpec[1].LIMIT = 4000;
    m_tierSpec[2].LIMIT = 1000;
    m_tierSpec[3].LIMIT = 500;
  }

  /**
   * Halo admin modifier
   *
   */
  modifier onlyAdmin() {
    require(msg.sender == m_addrAdmin);
    _;
  }

  /**
   * Halo MNAdmin modifier
   *
   */
  modifier onlyMNAdmin() {
    require(msg.sender == m_mnAdminAddr);
    _;
  }

  /**
   * MNOwnership SC modifier
   *
   */
  // modifier onlyMNOwnershipProposed() {
  //   // When a node is not existing, id = 0
  //   require(m_nodeMap[msg.sender].id > 0);
  //   _;
  // }

  /**
   * Remove a proposed MN out of list for MNAdmin
   *
   * @param _nodeAddr  MNInstance address
   */
  function adminRemoveMN(address _nodeAddr) external onlyMNAdmin {

    uint index;
    uint length = m_nodeAddrList.length;
    uint8 tier;
    for (index = 0; index < length; index++){
        if (m_nodeAddrList[index] == _nodeAddr) {
          break;
        }
    }

    /// There is no use-case for inexistent
    if(index < length) {
        /// Remove out of list and map
      if ( index + 1 != length ) {
        /// Move the last to the current
        m_nodeAddrList[index] = m_nodeAddrList[length - 1];
      }
      delete m_nodeAddrList[length - 1];
      m_nodeAddrList.length--;
    }
    /// Decrease limit

    (,tier,,) = getNodeSumInfo(_nodeAddr);
    if (m_tierSpec[tier-1].count > 0) {
      m_tierSpec[tier-1].count--;
    }
  }

  /**
   * Remove a proposed MN out of list
   *
   * @param _nodeAddr  MNOwnershipProposed SC address
   */
  function removeMN(address _nodeAddr) private {

    uint index;
    uint length = m_nodeAddrList.length;
    uint8 tier;
    for (index = 0; index < length; index++){
        if (m_nodeAddrList[index] == _nodeAddr) {
          break;
        }
    }

    /// There is no use-case for inexistent
    if(index < length) {
        /// Remove out of list and map
      if ( index + 1 != length ) {
        /// Move the last to the current
        m_nodeAddrList[index] = m_nodeAddrList[length - 1];
      }
      delete m_nodeAddrList[length - 1];
      m_nodeAddrList.length--;
    }
    /// Decrease limit

    (,tier,,) = getNodeSumInfo(_nodeAddr);
    if (m_tierSpec[tier-1].count > 0) {
      m_tierSpec[tier-1].count--;
    }
    // delete m_nodeMap[_nodeAddr];
  }


  /**
   * Remove a proposed MN out of list
   *
   * @param _nodeAddr  MNOwnershipProposed SC address
   */
  function addNode(NodeInfo _node_info, address[] _other_owners, uint[] _other_shares, string _ipv4_address, string _alias) private returns (address _nodeAddr, uint32 _id){
    // Put into MN list
    _node_info.id = ++m_ui_idGen;
    address sc_nodeAddr = (new MNInstance).value(msg.value)(address(this), msg.sender, _other_owners, _other_shares);
//    m_nodeMap[sc_nodeAddr] = _node_info;
    m_nodeAddrList.push(sc_nodeAddr);
    m_tierSpec[_node_info.tier - 1].count++;

    MNInstance mn_instance = MNInstance(sc_nodeAddr);
    mn_instance.setBasicInfo(_node_info.tier, m_tierSpec[_node_info.tier - 1].COST, _node_info.is_private, _node_info.id,
                               _node_info.state, _node_info.is_halo_hosting, _ipv4_address, _alias);

   mn_instance.setSecurityAddrs(m_mnAdminAddr, m_mnMarketplaceAddr, m_mnRewardAddr);
    return (sc_nodeAddr, _node_info.id);
  }

  /**
   * Internal checking for share
   * 
   * @param _node_info          Node information
   * @param _other_joiners      Joiner of private party
   * @param _other_shares       Share of each joiner
   *
   */
  function checkCondition(NodeInfo _node_info, address[] _other_joiners, uint[] _other_shares) 
                              private view returns (uint8 _state) { 
    uint8 state;
    // tierType must be [1, 4]
    require( _node_info.tier >= 1 && _node_info.tier <= 4);
    // Halo cost checking
    require( m_tierSpec[_node_info.tier - 1].count < m_tierSpec[_node_info.tier - 1].LIMIT);

     /// Public node
    if ( _node_info.is_private == false ) {
      /// restricted to tier1 and tier2, 
      require( _node_info.tier == 1 || _node_info.tier == 2);
    }

    // Check deposit: minimum for deposit and full ether
    // require( msg.value >= 1 ether)
    require( msg.value >= MIN_HALO_SHARE_K && msg.value <= m_tierSpec[_node_info.tier - 1].COST && ((msg.value % 1 ether) == 0));

    /// Public node
    if ( _node_info.is_private == false ) {
      /// must be hosted by Halo platform
      require( _node_info.is_halo_hosting == true);
      /// Multi-owner without pre-setting a list of co-owners
      require( _other_joiners.length == 0 &&  msg.value < m_tierSpec[_node_info.tier - 1].COST);
      state = MN_STATE_INITIALIZING_K; /// Intialize a shared MN
    }
    else {
      /// Multi-owner
      if (msg.value < m_tierSpec[_node_info.tier - 1].COST) {
        /// Check the share configuration
        require(_other_joiners.length > 0 && _other_shares.length == _other_joiners.length);
        /// Check whether a share is full Halo
        uint total_share = msg.value;
        for (uint i = 0; i < _other_joiners.length; i++) {
          require( _other_shares[i] >= MIN_HALO_SHARE_K &&((_other_shares[i] % 1 ether) == 0));
          total_share += _other_shares[i];
        }

        // total share equals to COST
        require(total_share == m_tierSpec[_node_info.tier - 1].COST);
        state = MN_STATE_INITIALIZING_K; /// Intialize a shared MN
      }
      else {
        /// Sole owner
        state = MN_STATE_DEPOSITED_K;    /// Single owner node  
      }
    }

    return state;
  }

  /**
   * Set MNAdminAddr for security. 
   * 
   * @param _mnAdminAddr          MNAdmin address
   *
   */
  function setMNAdminAddr(address _mnAdminAddr) external onlyAdmin {
    require(_mnAdminAddr > address(0));
    m_mnAdminAddr = _mnAdminAddr;
  }

  /**
   * Set MNMarketplace for security. 
   * 
   * @param _mnMarketplaceAddr          MNMarketplace address
   *
   */
  function setMNMarketplaceAddr(address _mnMarketplaceAddr) external onlyAdmin {
    require(_mnMarketplaceAddr > address(0));
    m_mnMarketplaceAddr = _mnMarketplaceAddr;
  }

  /**
   * Set MNReward for security. 
   * 
   * @param _mnRewardAddr          MNReward address
   *
   */
  function setMNRewardAddr(address _mnRewardAddr) external onlyAdmin {
    require(_mnRewardAddr > address(0));
    m_mnRewardAddr = _mnRewardAddr;
  }

  /**
   * Initialize a proposed MN. If the deposit is full, automatically apply
   * 
   * @param _tier               Tier type (1 -> 4)
   * @param _is_private         Node is private
   * @param _is_halo_hosting    MN is hosted by Halo platform
   * @param _other_owners       In case of private multi-owners MN, the addresses of other co-owners must be pre-set.
   * @param _other_shares       In case of private multi-owners MN, the corresponding share each co-owner must be pre-set. Value is Halo unit
   *
   * @return Event_InitializeSharedMN  A shared MN, state=INITIALIZING
   * @return Event_ApplyMN             A single owner MN, state=DEPOSITED
   *
   */
  function initializeMN(uint8 _tier, bool _is_private, bool _is_halo_hosting, string _alias, string _ipv4_address,
                          address[] _other_owners, uint[] _other_shares) external payable {
    /// Check the conditions
    /// - Sole Owner -> Private -> HALO/SELF Hosted
    /// - Multi Owner -> Private -> HALO/SELF Hosted
    /// - Multi Owner -> Public -> HALO Hosted
    /// - Public node are restricted to tier1 and tier2
    NodeInfo memory node_s;
    address sc_nodeAddr;

    require(m_mnAdminAddr > address(0));
    require(m_mnMarketplaceAddr > address(0));
    require(m_mnRewardAddr > address(0));

    node_s.tier = _tier;
    node_s.is_private = _is_private;
    node_s.is_halo_hosting = _is_halo_hosting;
    node_s.state = checkCondition(node_s, _other_owners, _other_shares);
    (sc_nodeAddr, node_s.id) = addNode(node_s, _other_owners, _other_shares, _ipv4_address, _alias);

    /// Raise corresponding event
    if (node_s.state == MN_STATE_INITIALIZING_K)
      Event_MNFactory(sc_nodeAddr, node_s.id, MN_EVT_InitializeMN);
    else
      Event_MNFactory(sc_nodeAddr, node_s.id, MN_EVT_FinalizeMN);
  }

  /**
   * Destroy MN
   * 
   * @param _nodeAddr        MNInstance address
   *
   *
   */
  function destroyMN(address destroyer, address _nodeAddr) private {
    MNInstance mn_instance = MNInstance(_nodeAddr);
    mn_instance.destroy(destroyer);
  }

  /**
   * Destroy MN for MNAdmin
   * 
   * @param _nodeAddr        MNInstance address
   *
   *
   */
  function adminDestroyMN(address destroyer, address _nodeAddr) public onlyMNAdmin {
    MNInstance mn_instance = MNInstance(_nodeAddr);
    mn_instance.destroy(destroyer);
  }

  /**
   * Join to a shared MN, become a co-owner
   * 
   * @param _nodeAddr        MNInstance address
   *
   *
   */
  function joinToInitializedMN(address _nodeAddr) external payable {
    MNInstance mn_instance = MNInstance(_nodeAddr);
    mn_instance.join.value(msg.value)(msg.sender);
    uint32 nodeId;
    uint8 nodeState;

    (,,nodeId, nodeState,,,) = mn_instance.getBasicInfo();
    if (nodeState == MN_STATE_DEPOSITED_K) {
      Event_MNFactory(_nodeAddr, nodeId, MN_EVT_FinalizeMN);
    }
  }

    /**
   * Withdraw share
   * 
   * @param _nodeAddr        MNInstance address
   * @param _withdrawValue   Share to withdraw
   *
   *
   */
  function withdrawFromInitializedMN(address _nodeAddr, uint _withdrawValue) external {
    MNInstance mn_instance = MNInstance(_nodeAddr);
    mn_instance.withdraw(msg.sender, _withdrawValue);
    uint32 nodeId;
    uint8 nodeState;

    (,,nodeId, nodeState,,,) = mn_instance.getBasicInfo();
    if (nodeState == MN_STATE_INVALID_K) {
      removeMN(_nodeAddr);
      destroyMN(msg.sender, _nodeAddr);
      Event_MNFactory(_nodeAddr, nodeId, MN_EVT_AutoCancel);
    }
  }

  function getNodeList() external view returns (address[] _node_list) {
    return m_nodeAddrList;
  }

  function getNodeSumInfo(address _nodeAddr) public view 
        returns (uint32 _id, uint8 _tier, uint8 _state, bool _is_private) {
    MNInstance mn_instance = MNInstance(_nodeAddr);
    uint32 nodeId;
    uint8 nodeTier;
    uint8 nodeState;
    bool isPrivate;

    (nodeTier, isPrivate, nodeId, nodeState,,,) = mn_instance.getBasicInfo();
    return (nodeId, nodeTier, nodeState, isPrivate);
  }

  function getNodeCount(uint8 _tier) external view returns(uint16 _count) {
    return (m_tierSpec[_tier-1].count);
  }

  function getActivatedNodeCount(uint8 _tier) external view returns(uint16 _count) {
    return (m_tierSpec[_tier-1].activatedCount);
  }

  function increaseActivatedCount(uint8 _tier) external {
    m_tierSpec[_tier-1].activatedCount++;
  }

  function getMNCount() external view returns(uint _count) {
    return m_nodeAddrList.length;
  }

  function getMNAddrByIndex(uint _index) external view returns(address _mnAddr) {
    return m_nodeAddrList[_index];
  }
}
