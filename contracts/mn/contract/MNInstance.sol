pragma solidity ^0.4.18;

/** @title SC for each MN. */
contract MNInstance {
  /// Set minimum HALO of a share. Prevent from creating small share. The owner list become very long.
  uint private constant MIN_HALO_SHARE_K = 1 ether;
  /// Node state
  uint8 private constant MN_STATE_INVALID_K       = 0;
  uint8 private constant MN_STATE_INITIALIZING_K  = 1;
  uint8 private constant MN_STATE_DEPOSITED_K     = 2;
  uint8 private constant MN_STATE_ACTIVE_K        = 3;
  uint8 private constant MN_STATE_SUSPEND_K       = 4;

  uint8 private constant EVENT_TYPE_JOIN_SHARE_K  = 1; /// Event when join successfully
  uint8 private constant EVENT_TYPE_WITHDRAW_K    = 2; /// Event when witdraw a share
  uint8 private constant EVENT_TYPE_APPLY_K       = 3; /// Event when apply a MN to Halo admin
  uint8 private constant EVENT_TYPE_DESTROY_K     = 4; /// Event when Halo admin destroy MN
  uint8 private constant EVENT_TYPE_PROMOTE_K     = 5; /// Event when promoting to actual MNOwnership

  uint8             private   m_state;
  uint32            private   m_id;
  address           private   m_factoryAddr;
  address           private   m_adminAddr;
  address           private   m_marketplaceAddr;
  address           private   m_rewardAddr;
//  MNProposedMarket  private  m_factorySC;
  uint              private   m_depositedHalo;
  uint              private   m_cost;
  uint8             private   m_tier;
  bool              private   m_isHaloHosting;
  string            private   m_ipv4Address;
  string            private   m_alias;
  bool              private   m_isPrivate;
  uint              private   m_firstPing;
  uint              private   m_lastPing;
  address           private   m_pingClient;


  address[]                   private m_joinerAddrList;         /// Joiner who deposited
  mapping (address => uint)   private m_joinerShareMap;               /// Map (joiner address -> deposit share in HALO)
  address[]                   private m_privateJoinerAddrList; /// Private joiner who is added by 1st joiner when private party
  mapping (address => uint)   private m_privateJoinerShareMap;       /// Map (private joiner address -> required deposit share in HALO)
  mapping(address => uint)    private m_shareLockMap;   /// Lock owner's share when it's in selling list

  /// Event
  event Event_Action(address indexed _joinerAddr, address indexed _nodeAddr, uint32 _nodeId, uint8 _eventType);
  event Event_AdminChangeState(uint8 _state);
  event Event_Reward(uint32 _nodeId, uint _value);

  /**
   * Halo MNFactory modifier
   *
   */
  modifier onlyMNFactory() {
    require(msg.sender == m_factoryAddr);
    _;
  }

  /**
   * Halo MNAdmin modifier
   *
   */
  modifier onlyMNAdmin() {
    require(msg.sender == m_adminAddr);
    _;
  }

  /**
   * Halo MNMarketplace modifier
   *
   */
  modifier onlyMNMarketplace() {
    require(msg.sender == m_marketplaceAddr);
    _;
  }

  /**
   * Halo MNReward modifier
   *
   */
  modifier onlyMNReward() {
    require(msg.sender == m_rewardAddr);
    _;
  }

  /**
   * Only joiner with his share > 0
   *
   */
  // modifier onlyJoiner() {
  //   require(m_joinerShareMap[msg.sender] > 0);
  //   _;
  // }

  /**
   * Constructor
   *
   * @param _factoryAddr      MNProposedMarket SC address, only use the address to interact with MNProposedMarket
   * @param _firstJoiner           The address of first joiner who initialized the MN, msg.value is the deposit of 1st joiner
   * @param _otherJoiners          The address of first owner who initialized the MN
   * @param _otherShares           The fixed share of each private joiner
   *
   */
  function MNInstance(address _factoryAddr, address _firstJoiner, address[] _otherJoiners, uint[] _otherShares) public payable {
    /// Set input
    m_factoryAddr = _factoryAddr;    /// Set Halo Admin adress
//    m_factorySC = MNProposedMarket(m_factoryAddr);

    m_depositedHalo = msg.value;
    /// Store joiner and private-joiner if having
    m_joinerShareMap[_firstJoiner] = msg.value;
    m_joinerAddrList.push(_firstJoiner);
    if (_otherJoiners.length > 0 && _otherShares.length == _otherJoiners.length ) {
      for ( uint i = 0; i < _otherJoiners.length; i++ ) {
        m_privateJoinerAddrList.push(_otherJoiners[i]);
        m_privateJoinerShareMap[_otherJoiners[i]] = _otherShares[i];
      }
    }
  }

  /**
   * 
   *
   * @param _state      State                  
   *
   */
  function setState(uint8 _state) external onlyMNAdmin {
    m_state = _state;
    Event_AdminChangeState(m_state);
  }

  /**
   * 
   *
   * @param _firstPing      First Ping
   * @param _lastPing       Last Ping
   *
   */
  function setPing(uint _firstPing, uint _lastPing) external onlyMNAdmin {
    m_firstPing = _firstPing;
    m_lastPing = _lastPing;
  }

  /**
   * 
   *
   * @param _state                  MN state
   *
   */
  function getState() public view returns(uint8 _state) {
    return (m_state);
  }

    /**
   * 
   *
   * @param _id                  MN id
   *
   */
  function getId() public view returns(uint32 _id) {
    return (m_id);
  }

    /**
   * 
   *
   * @param _tier                  MN id
   *
   */
  function getTier() public view returns(uint8 _tier) {
    return (m_tier);
  }
  

  /**
   * SetBasicInfo
   *
   * @param _tier                   Tier type (1->4)
   * @param _tierPrice              Tier price to check maximum deposit
   * @param _isPrivate              If MN is private, it's not showed in Marketplace
   * @param _proposedId             Human-readable ID
   * @param _state                  State of node: DEPOSITED(full deposit)/INITIALIZING(insufficient deposit)
   * @param _isHaloHosting          MN is hosted by Halo platform
   * @param _ipv4Address            In case of self-hosting, IPv4 address will be required
   * @param _alias                  Alias name
   *
   */
  function setBasicInfo(uint8 _tier, uint _tierPrice, bool _isPrivate, uint32 _proposedId, 
                        uint8 _state,  bool _isHaloHosting, string _ipv4Address, string _alias) external onlyMNFactory {
    /// Set input
    m_tier = _tier;
    m_cost = _tierPrice; /// set cost of the MasterNode
    m_isPrivate = _isPrivate;

    m_id = _proposedId;
    m_state = _state;
    m_isHaloHosting = _isHaloHosting;
    m_ipv4Address = _ipv4Address;
    m_alias = _alias;
  }

    /**
   * setAdvancedInfo
   *
   * @param _state                  State
   * @param _firstPing              firstPing
   * @param _lastPing               lastPing
   * @param _pingClient             Ping client Address
   *
   */
  function setAdvancedInfo(uint8 _state, uint _firstPing, uint _lastPing, address _pingClient, string _ipv4Address) external onlyMNAdmin {
    /// Set input
    m_state = _state;
    m_firstPing = _firstPing;
    m_lastPing = _lastPing;
    m_pingClient = _pingClient;
    m_ipv4Address = _ipv4Address;
  }

  /**
   * Get a advanced information
   * 
   * @return _state                   State
   * @return _firstPing               FirstPing
   * @return _lastPing                LastPing
   * @return _pingClient              PingClient
   *
   */
  function getAdvancedInfo() external view 
                returns(uint8 _state, uint _firstPing, uint _lastPing, address _pingClient, string _ipv4Address){
    return (m_state, m_firstPing, m_lastPing, m_pingClient, m_ipv4Address);
  }

  /**
   * Remove a joinner out of list
   *
   * @param _joinerAddr  Joinner address
   */
  function removeJoiner(address _joinerAddr) private {
    uint index;
    uint length = m_joinerAddrList.length;
    for (index = 0; index < length; index++){
        if (m_joinerAddrList[index] == _joinerAddr) {
          break;
        }
    }

    /// There is no use-case for inexistent
    assert(index < length);
    /// Remove out of list and map
    if ( index + 1 != length ) {
      /// Move the last to the current
      m_joinerAddrList[index] = m_joinerAddrList[length - 1];
    }
    delete m_joinerAddrList[length - 1];
    m_joinerAddrList.length--;
    delete m_joinerShareMap[_joinerAddr];
  }

  /**
   * Join to a shared MN, become a co-owner
   *
   */
  function join(address _joiner) external payable onlyMNFactory{
    /// Only when MN is being initialized (state=INITIALIZING)
    require(m_state == MN_STATE_INITIALIZING_K);
    /// Check the Halo share
    require(msg.value >= MIN_HALO_SHARE_K && ((msg.value % 1 ether) == 0));
    require(m_depositedHalo + msg.value <= m_cost);                  /// Not allow deposit over maximum price
    uint remaining_share = m_cost - m_depositedHalo - msg.value;     /// Always >= 0
    require(remaining_share == 0 || remaining_share >= MIN_HALO_SHARE_K );  /// Ensure the rest share >= MIN_HALO_SHARE_K or zero

    /// Public party
    if (m_isPrivate == false) {
      /// Increase the share of existing user
      if (m_joinerShareMap[_joiner] > 0) {
        m_joinerShareMap[_joiner] += msg.value;
      }
      else {
        /// Set a new co-owner
        m_joinerShareMap[_joiner] = msg.value;
        m_joinerAddrList.push(_joiner);
      }
    }
    else {
      /// Private party

      /// msg.value > MIN_HALO_SHARE_K is aready checked 
      /// Check whether sender is private joiner
      require(m_privateJoinerShareMap[_joiner] == msg.value);
      /// Not double-joining
      require(m_joinerShareMap[_joiner] == 0);

      /// Add to joiner list
      m_joinerShareMap[_joiner] = msg.value;
      m_joinerAddrList.push(_joiner);
    }
    m_depositedHalo += msg.value;

    /// When fund is enough, apply to Halo admin
    if (m_depositedHalo == m_cost) {
      /// Change state, not allowing to withdraw
      m_state = MN_STATE_DEPOSITED_K;
      /// Raise event
      Event_Action(_joiner, this, m_id, EVENT_TYPE_JOIN_SHARE_K);
      Event_Action(_joiner, this, m_id, EVENT_TYPE_APPLY_K);
      return;
    }
    /// Calling the event after ApplyMN because applying can be reverted
    Event_Action(_joiner, this, m_id, EVENT_TYPE_JOIN_SHARE_K);
  }

  /**
   * Withdraw a share
   *
   * @param _withdrawValue   A share value in Halo
   *
   */
  function withdraw(address withdrawer, uint _withdrawValue) external onlyMNFactory{
    /// Only when MN is being initialized (state=INITIALIZING)
    require(m_state == MN_STATE_INITIALIZING_K);

    require(m_joinerShareMap[withdrawer] > 0);
    /// Only allow withdrawing completely
    require(_withdrawValue == m_joinerShareMap[withdrawer]);

    /// Remove joiner
    removeJoiner(withdrawer);

    /// Send back HALO
    (withdrawer).transfer(_withdrawValue);
    m_depositedHalo -= _withdrawValue;
        
    /// Successfully
    Event_Action(withdrawer, this, m_id, EVENT_TYPE_WITHDRAW_K);

    if (m_depositedHalo == 0) {
      m_state = MN_STATE_INVALID_K;
    }
  }

  /**
   * Destroy the SC from Halo admin
   *
   */
  function destroy(address withdrawer) external onlyMNFactory {
    uint index;
    uint length = m_joinerAddrList.length;

    /// Send back HALO
    for (index = 0; index < length; index++){
      uint tmp_share = m_joinerShareMap[m_joinerAddrList[index]];
      /// Set balance to zero before transfering
      m_joinerShareMap[m_joinerAddrList[index]] = 0;
      (m_joinerAddrList[index]).transfer(tmp_share);
      m_depositedHalo -= tmp_share; 
    }
    //assert(m_depositedHalo == 0);

    Event_Action(withdrawer, this, m_id, EVENT_TYPE_DESTROY_K);
    /// Destroy
    selfdestruct(withdrawer);
  }

  /**
   * Get a basic information
   * 
   * @return _tier                   Tier type (1->4)
   * @return _isPrivate             If MN is private, it's not showed in Marketplace
   * @return _proposedId            Human-readable ID
   * @return _state                  State of node: DEPOSITED(full deposit)/INITIALIZING(insufficient deposit)
   * @return _ipv4Address           IPv4 address of MN host
   * @return _alias                  Alias name
   * @return _total_deposit          The current deposit
   *
   */
  function getBasicInfo() external view
                returns(uint8 _tier, bool _isPrivate, uint32 _proposedId, uint8 _state, 
                        string _ipv4Address, string _alias, uint _total_deposit){
    return (m_tier, m_isPrivate, m_id, m_state, m_ipv4Address, m_alias, m_depositedHalo);
  }

  /**
   * Get a joiner list
   * 
   */
  function getJoinerList() external view returns(address[] _joiner_list){

    return m_joinerAddrList;
  }

  /**
   * Get a joiner share
   * 
   * @param _joinerAddress       Joiner address
   *
   * @return _joinerAddress      Joiner's share
   *
   */
  function getJoinerShare(address _joinerAddress) external view returns(uint _joiner_share){

    return m_joinerShareMap[_joinerAddress];
  }

  /**
   * Get a private joiner list
   * 
   */
  function getPrivateList() external view returns(address[] _joiner_list){

    return m_privateJoinerAddrList;
  }

  /**
   * Get a private joiner share
   * 
   * @param _joinerAddress       Joiner address
   *
   * @return _joinerAddress      Joiner's share
   *
   */
  function getPrivateShare(address _joinerAddress) external view returns(uint _joiner_share){

    return m_privateJoinerShareMap[_joinerAddress];
  }

  function getInfo() external view
              returns (uint8 _state, uint32 _node_id, uint _total_deposit,
                      uint8 _tier, bool _is_halo_hosting, address _ping_client,
                      uint _first_ping, uint _last_ping, string _ipv4_address, string _alias) {
    return (m_state, m_id, m_depositedHalo,
            m_tier, m_isHaloHosting, m_pingClient,
            m_firstPing, m_lastPing, m_ipv4Address, m_alias);
  }

  /**
   * Get a owner list
   * 
   */
  function getOwnerList() external view returns(address[] _ownerList){

    return m_joinerAddrList;
  }

  /**
   * Get a owner share
   * 
   * @param _ownerAddress       Owner address
   *
   * @return _ownerShare      Owner's share
   *
   */
  function getOwnerShare(address _ownerAddress) external view returns(uint _ownerShare){

    return m_joinerShareMap[_ownerAddress];
  }

  /**
   * Get a owner share locked
   * 
   * @param _ownerAddress       Owner address
   *
   * @return _ownerShareLocked      Owner's share Locked
   *
   */
  function getShareLocked(address _ownerAddress) external view returns(uint _ownerShareLocked){

    return m_shareLockMap[_ownerAddress];
  }

    /**
   * Get a owner share remainig
   * 
   * @param _ownerAddress       Owner address
   *
   * @return _ownerShareUnlocked      Owner's share Unlocked
   *
   */
  function getShareUnlocked(address _ownerAddress) external view returns(uint _ownerShareUnlocked){

    return m_joinerShareMap[_ownerAddress] - m_shareLockMap[_ownerAddress];
  }

  /**
   * Lock owner's share for selling
   * 
   * @param _ownerAddress       Owner address
   * @param _shareToBeLocked          Owner share to be locked
   *
   */
  function lockShare(address _ownerAddress, uint _shareToBeLocked) external onlyMNMarketplace {
    if ( m_shareLockMap[_ownerAddress] == 0) {
      m_shareLockMap[_ownerAddress] = _shareToBeLocked;
    }
    else {
      m_shareLockMap[_ownerAddress] += _shareToBeLocked;
    }
  }

  /**
   * Lock owner's share for selling
   * 
   * @param _ownerAddress       Owner address
   * @param _shareToBeUnlocked          Owner share to be locked
   *
   */
  function unlockShare(address _ownerAddress, uint _shareToBeUnlocked) external onlyMNMarketplace {
    m_shareLockMap[_ownerAddress] -= _shareToBeUnlocked;
    if (m_shareLockMap[_ownerAddress] == 0) {
      delete m_shareLockMap[_ownerAddress];
    }
  }

  /**
   * Transfer ownership to buyer from seller
   * 
   * @param seller        seller address
   * @param buyer        buyer address
   * @param _shareSold        share sold
   *
   */
  function transferOwnership(address seller, address buyer, uint _shareSold) external payable onlyMNMarketplace {
    /// Transfer ownership
    m_joinerShareMap[seller] -= _shareSold;
    if (m_joinerShareMap[seller] == 0) {
      /// Remove owner if share=0
      removeJoiner(seller);
    }

    /// Increase share for an existing owner
    if (m_joinerShareMap[buyer] > 0) {
      m_joinerShareMap[buyer] += _shareSold;
    }
    else {
      /// Add a new owner
      m_joinerShareMap[buyer] = _shareSold;
      m_joinerAddrList.push(buyer);
    }

    /// Transfer money to seller
   (seller).transfer(msg.value);
  }

  /**
   * Distribute reward for each co-owner
   *
   */
  function distributeReward() external payable onlyMNReward {
    /// Transfer money to each co-owner
    
    uint total_reward = msg.value;

    /// Copy owner share list to prevent sell/buy a share in fallback function
    uint list_len = m_joinerAddrList.length;
    uint[] memory share_list = new uint[](list_len);
    address[] memory addr_list = new address[](list_len);
    uint i;
    for( i = 0; i < m_joinerAddrList.length; i++) {
      addr_list[i] = m_joinerAddrList[i];
      share_list[i] = m_joinerShareMap[m_joinerAddrList[i]];
    }

    uint temp;
    for(i = 0; i < addr_list.length; i++) {
      /// Last owner
      if ( (i + 1) == addr_list.length) {
        temp = total_reward;
      }
      else {
        temp = msg.value * share_list[i] / m_depositedHalo;
      }
      total_reward -= temp;
      (addr_list[i]).transfer(temp);
    }
    Event_Reward(m_id, msg.value);
  }

  /**
   * Ping to MN for update status
   * 
   * @param _pinger                   Pinger's address
   * @param _ipv4Address              IP address
   *
   */
  function doPing(address _pinger, string _ipv4Address) external onlyMNReward {
    /// Only for ACTIVE MN
    require(m_state == MN_STATE_ACTIVE_K);
    require(m_pingClient != address(0) && _pinger == m_pingClient);

    /// Check IP address
    require(keccak256(m_ipv4Address) == keccak256(_ipv4Address));

    // Update PING time
    if ( m_firstPing == 0 ){
      m_firstPing = now;
    }
    m_lastPing = now;
  }

  /**
   * Set Addresses for security
   * 
   * @param _mnAdmin                    MNAdmin address
   * @param _mnMarketplace              MNMarketplace address
   * @param _mnReward                   MNReward address
   *
   */
  function setSecurityAddrs(address _mnAdmin, address _mnMarketplace, address _mnReward) external onlyMNFactory{
    /// Only for ACTIVE MN
    require(_mnAdmin > address(0));
    require(_mnMarketplace > address(0));
    require(_mnReward > address(0));

    m_adminAddr = _mnAdmin;
    m_marketplaceAddr = _mnMarketplace;
    m_rewardAddr = _mnReward;
  }

  /**
   * Get Scurity Addresses
   *
   * @return _factory       Factory Address
   * @return _admin         Admin Address
   * @return marketplace    Marketplace Address
   * @return reward         Reward Address
   *
   */
  function getSecurityAddrs() external view returns(address _factory, address _admin, address marketplace, address reward){

    return (m_factoryAddr, m_adminAddr, m_marketplaceAddr, m_rewardAddr);
  }  
}
