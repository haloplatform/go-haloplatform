pragma solidity ^0.4.18;

import "./MNFactory.sol";

/** @title MN management SC for Halo admin. */
contract MNReward {
  /// Set minimum HALO of a share
  uint private constant MIN_HALO_SHARE_K = 1 ether;

  uint private constant HALO_PoW_REWARD       = 28 ether;
  uint private constant HALO_CORE_PoW_REWARD  = 0.5 ether;
  uint private constant HALO_NODE_PoW_REWARD  = 27.5 ether;

  uint8 private constant SFP_DIVISION_PC_HALO_CORE = 35;   /// 35% for Halo core
  uint8 private constant SFP_DIVISION_PC_MN        = 50;   /// 50% for MNs
  uint8 private constant SFP_DIVISION_PC_MINER     = 15;   /// 15% for miner

  /// Node state
  uint8 private constant MN_STATE_INVALID_K       = 0;
  uint8 private constant MN_STATE_ACTIVE_K        = 3;
  uint8 private constant MN_STATE_SUSPEND_K       = 4;

  uint8 private constant MN_EVT_MNCreate          = 1;
  uint8 private constant MN_EVT_MNSuspend         = 2;
  uint8 private constant MN_EVT_MNResume          = 3;
  uint8 private constant MN_EVT_MNCancel          = 4;

  /// We can configure shorter or longer for testing
  uint32 private constant SFP_DISTRIBUTE_INTEVAL   = 1 days;                               /// Daily=86400s
  uint32 private constant MAX_PoW_PER_PERIOD       = SFP_DISTRIBUTE_INTEVAL / (4 minutes); /// PoW every 4 minutes: val=360

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

  /// Management storage
  address     private  m_addr_admin;
  address     private  m_halo_fund_addr;   /// Address of Halo fund, initial_value=m_addr_admin 
  address     private  m_proposed_mkt_addr;
  uint32      private  m_ui_nodeIDGen;
  TierInfo[4] private  m_tierSpec;
  address[]                     private  m_node_addr_list;
  mapping (address => NodeInfo) private  m_node_map;

  /// Service fee
  uint16 private constant   MN_DIE_UPTIME     = 4 minutes;  /// If a node don't update PING after 4 minutes, it's considered as die.
  uint32 private constant   MN_REWARD_UPTIME  = 1 days;     /// At least 24h uptime for rewarding
  mapping (uint8 => address[])   private    m_pow_list;     /// A list of MN for PoW distribution
  uint      private   m_next_sfp;                   /// Service fee pool(SFP) for next day distribution
  uint      private    m_miner_sfp_remaining;       /// the remaing miner SFP for today, distribution to miner every PoW
  uint      private    m_miner_sfp_share;           /// the miner SF reward value for each miner: value=m_sfp/MAX_PoW_PER_PERIOD; 
  uint16    private    m_miner_pow_cnt;             /// The counter for distributing service fee reward to miner: <= MAX_PoW_PER_PERIOD
  uint      private    m_miner_last_block;          /// Trace this to prevent reentrancy
  uint      private    m_time_last_distribute;      /// We will distribute SF daily(SFP_DISTRIBUTE_INTEVAL)
  address   private    m_mnFactoryAddr;

  /// Event
  event Event_MNReward(address indexed _mn_instanceship_addr, uint32 _node_id, uint8 _event_type);

  /**
   * Constructor
   */
  function MNReward(address _mnFactoryAddr) public {
    // constructor
    m_addr_admin = msg.sender;
    m_halo_fund_addr = m_addr_admin;
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
    /// Limitation
    m_tierSpec[0].PoW = 4.5 ether;
    m_tierSpec[1].PoW = 7.5 ether;
    m_tierSpec[2].PoW = 5.5 ether;
    m_tierSpec[3].PoW = 10 ether;

    /// Counting from deploying time
    m_time_last_distribute = now;
    m_miner_last_block = block.number;
    m_mnFactoryAddr = _mnFactoryAddr;
  }

  /**
   * Halo admin modifier
   *
   */
  modifier onlyAdmin() {
    require(msg.sender == m_addr_admin);
    _;
  }

//   /**
//    * A MN will call this API to notify its status
//    * 
//    */
//   function MNIsAlive() external onlyMNOwnership {
//     /// Only for ACTIVE MN
//     require(m_node_map[msg.sender].state == MN_STATE_ACTIVE_K);

//     // Update PING time
//     if ( m_node_map[msg.sender].first_ping == 0 ){
//       m_node_map[msg.sender].first_ping = now;
//     }
//     m_node_map[msg.sender].last_ping = now;
//   }

/********************************** Reward distribution **********************************/
  /**
   * Send PoW rewards for 4 nodes in 4 tiers 
   * 
   */
  function PoWDistribution() internal {
    uint8    tier_ind;
    address[4] memory selected_nodes;

    /// Select only one node for each tier
    for (tier_ind = 0; tier_ind < 4; tier_ind++) {
      uint list_len = m_pow_list[tier_ind].length;
      uint i = 0;
      uint32  destroy_cnt = 0;
      /// Check for each tier until getting a node
      for ( i = 0; i < list_len && selected_nodes[tier_ind] == 0; i++) {
        address node_addr = m_pow_list[tier_ind][i];
        MNInstance mn_instance = MNInstance(node_addr);
        uint8 nodeState;
        uint firstPing;
        uint lastPing;
        (nodeState, firstPing, lastPing,,) = mn_instance.getAdvancedInfo();
        if(nodeState == MN_STATE_ACTIVE_K 
          && firstPing > 0 && now - lastPing > MN_DIE_UPTIME
          && lastPing - firstPing >= MN_REWARD_UPTIME) {
          /// Select node
          selected_nodes[tier_ind] = node_addr;
        }

        /// Node is destroyed, not in map 
        /// lazy deleting: node is not removed out of this list when MNCancel is called
        if (nodeState == MN_STATE_INVALID_K) {
          destroy_cnt++;
        }
        else {
          /// Push back into the end of list. We will shift later.
          m_pow_list[tier_ind].push(node_addr);
        }
      }

      /// i is distance which we will shift
      uint node_ind = 0;
      for (node_ind = 0; node_ind < list_len - destroy_cnt; node_ind++) {
         m_pow_list[tier_ind][node_ind] = m_pow_list[tier_ind][node_ind + i];
      }
      /// Reset the length
      m_pow_list[tier_ind].length = list_len - destroy_cnt;
    }

    /// Transfer reward to selected node
    for (tier_ind = 0; tier_ind < 4; tier_ind) {
      if ( selected_nodes[tier_ind] > 0 ) {
        mn_instance = MNInstance(selected_nodes[tier_ind]);
        mn_instance.distributeReward.value(m_tierSpec[tier_ind].PoW)();
      }
      else {
        /// If there is no node, push the reward to SFP
        m_next_sfp += m_tierSpec[tier_ind].PoW;
      }
    }

    /// Transfer 0.5 Halo to Halo core
    (m_halo_fund_addr).transfer(HALO_CORE_PoW_REWARD);
  }

  function tierDistribute(uint[4] _tier_share) private returns (uint _remaining) {
    uint[4] memory tier_node_reward;
    uint total_reward = 0;
    uint index;
    MNFactory mn_factory = MNFactory(m_mnFactoryAddr);
    /// Calculate the reward for each MN, based on tier
    for (index = 0; index < 4; index++) {
      if ( m_tierSpec[index].count == 0) {
        /// The share of the tier, which have no node, will be rolled to SFP for next day
        tier_node_reward[index] = 0;
      }
      else {
        /// Not need to count the unqualified nodes (Ex: uptime < 24h or SUSPEND)
        /// The reward for unqualified nodes will be rolled to SFP for next day
        tier_node_reward[index] = _tier_share[index] / m_tierSpec[index].activatedCount;
      }
      /// For returning the remaining
      total_reward += _tier_share[index];
    }

    for ( index = 0; index < mn_factory.getMNCount() && total_reward > 0; index++ ) {
      address nodeAddr = mn_factory.getMNAddrByIndex(index);
      MNInstance mn_instance = MNInstance(nodeAddr);
      uint8 nodeState;
      uint firstPing;
      uint lastPing;
      uint8 nodeTier;
      (nodeState, firstPing, lastPing,,) = mn_instance.getAdvancedInfo();
      (nodeTier,,,,,,) = mn_instance.getBasicInfo();

      if ( nodeState == MN_STATE_ACTIVE_K 
          && lastPing > 0 && now - lastPing > MN_DIE_UPTIME
          && lastPing - firstPing >= MN_REWARD_UPTIME) {
        /// Subtract from total
        total_reward -= tier_node_reward[nodeTier];
        /// Node is active over 24h and still alive
        mn_instance = MNInstance(nodeAddr);
        mn_instance.distributeReward.value(tier_node_reward[nodeTier])();
      }
    }

    /// The remaining
    return total_reward;
  }

  function SFDistribution() private {
    /// Distributing reward for MN
    /// Roll the remaining of miner SFP
    uint total_sfp = m_miner_sfp_remaining + m_next_sfp;
    /// Reset SFP
    m_next_sfp = 0;

    /// Distribute SFP to: Halo core(35%), MNs(50%) and Miner SFP(15% for next day)
    /// Set the share for miner SFP
    m_miner_sfp_remaining = total_sfp * SFP_DIVISION_PC_MINER / 100;  /// Total reward for miner
    m_miner_sfp_share = m_miner_sfp_remaining / MAX_PoW_PER_PERIOD;  /// reward for each PoW

    /// Reward for MNs
    uint[4] memory tier_share;
    tier_share[0] = total_sfp * 50 / 1000;  /// 5%
    tier_share[1] = total_sfp * 75 / 1000;  /// 7.5%
    tier_share[2] = total_sfp * 125 / 1000; /// 12.5%
    tier_share[3] = total_sfp * 250 / 1000; /// 25%
    
    uint mn_part = tier_share[0] + tier_share[1] + tier_share[2] + tier_share[3];
    /// Roll the remaing to SFP
    m_next_sfp += tierDistribute(tier_share);

    assert(total_sfp > (m_miner_sfp_remaining + mn_part));
    /// Reward for Halo core
    (m_halo_fund_addr).transfer(total_sfp - m_miner_sfp_remaining - mn_part);
  }

  /**
   * Receive fee explicitly from all network.(not use fallback function)
   *
   */
  function receiveFee() external payable {
    /// Not allow zero
    require(msg.value > 0);

    /// Add to SFP for next day
    m_next_sfp += msg.value;
  }

  /**
   * Send PoW reward (28 Halo) to Halo Core and MNs
   *
   * @param  _miner_address       Miner address
   */
  function blockRewardMN(address _miner_address) external payable {
    /// Make sure 28 Halo
    /// TODO: more checking about calling from miner
    require(msg.value == HALO_PoW_REWARD);

    /// Transfer service fee reward to miner
    /// If the number of PoW blocks are mined over 360 that day, SF for miner will be 0. 
    if (m_miner_sfp_remaining > 0 && m_miner_sfp_remaining >= m_miner_sfp_share) {
      m_miner_sfp_remaining -= m_miner_sfp_share;
      (_miner_address).transfer(m_miner_sfp_share);
    }
    /// Distribute 28 Halo from PoW to MNs and Halo admin
    PoWDistribution();

    /// Check whether it's time to distribute for next period (daily)
    if(now <= m_time_last_distribute || now - m_time_last_distribute < SFP_DISTRIBUTE_INTEVAL ) {
      return;
    }
    /// Distributing SF to MNs (SFP_DISTRIBUTE_INTEVAL=24h)
    SFDistribution();
    /// Set timing for next distribution
    m_time_last_distribute = now;
  }

  /**
   * Ping to MN for update status
   * 
   * @param _mnAddr                   MN address to Ping
   * @param _ipv4Addr              IP address
   *
   */
  function doPing(address _mnAddr, string _ipv4Addr) external {
    MNInstance mn_instance = MNInstance(_mnAddr);
    uint8 nodeState;
    uint firstPing;
    uint lastPing;
    (nodeState, firstPing, lastPing,,) = mn_instance.getAdvancedInfo();

    require(nodeState == MN_STATE_ACTIVE_K);
    mn_instance.doPing(msg.sender, _ipv4Addr);
  }

  /**
   * Set address of Halo fund
   *
   * @param  _halo_fund_addr       The address to send to Halo core
   */
  function AdminSetFundAddr(address _halo_fund_addr) external onlyAdmin{
    m_halo_fund_addr = _halo_fund_addr;
  }
}