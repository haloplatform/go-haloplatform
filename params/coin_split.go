package params

import (
	"github.com/ethereum/go-ethereum/common"
)

type AccountInfo struct {
	Address    common.Address /// Account address
	AddBalance string         /// Convert to big.Int when AddBalance()
}

// Account list to inflate for Mainnet
func MainnetCoinSplitAccount() []AccountInfo {
	return []AccountInfo{
		{common.HexToAddress("0x67b7301878bc38559e118e9958d1bb8b4eb6afb2"), "199000000000000000000"}, /// 199 Halo
		{common.HexToAddress("0x928e5528aaa3c19640836c802c2743a0346acc56"), "299000000000000000000"}, /// 299 Halo
	}
}

// Account list to inflate for testnet
func TestnetCoinSplitAccount() []AccountInfo {
	return []AccountInfo{
		{common.HexToAddress("0x022e37293b64304e44e5ac7aaacc9b8d1d3b5880"), "1339923000000000000000000"},             /// New 1339923 Halo, Old 1677 Halo
		{common.HexToAddress("0x037f9578824ef63d76e1e0189f5063dca0337a80"), "79101000000000000000000000000"},         /// New 79101000000 Halo, Old 99000000 Halo
		{common.HexToAddress("0x043014414645ae0af338effa0ce71696565845da"), "48019900000000000000000000"},            /// New 48019900 Halo, Old 60100 Halo
		{common.HexToAddress("0x09c479bd9c0618ae602afb9e1b00dd0b684907bf"), "3947060000000000000000000000"},          /// New 3947060000 Halo, Old 4940000 Halo
		{common.HexToAddress("0x0dfcff0a11a5bdea4ea4e73155827e58b6270e74"), "3196000000000000000000"},                /// New 3196 Halo, Old 4 Halo
		{common.HexToAddress("0x0efe86dfea8b10aa7e6c11850766498cc12c4120"), "679150000000000000000000"},              /// New 679150 Halo, Old 850 Halo
		{common.HexToAddress("0x0f67d1751b2ed1da5e6ff0fb10429c0fd099131c"), "495380000000000"},                       /// New 0.00049538 Halo, Old 0.00000062 Halo
		{common.HexToAddress("0x0ff1867b84be3c2e4e643c8baab15e5b6d4e0ad9"), "450159796000000000000000000"},           /// New 450159796 Halo, Old 563404 Halo
		{common.HexToAddress("0x10a26b3b8c8a9c5ae6fcff43b85daf4491fc755d"), "79868040000000000000000000000"},         /// New 79868040000 Halo, Old 99960000 Halo
		{common.HexToAddress("0x126a69d78689467d85afcef481ae970bd61b86bc"), "357496370250000000000000"},              /// New 357496.37025 Halo, Old 447.42975 Halo
		{common.HexToAddress("0x17eb137ba4af604ddd747086a70b133b38bbc62b"), "7940462000000000000000000000"},          /// New 7940462000 Halo, Old 9938000 Halo
		{common.HexToAddress("0x1c77063ce0adae51091d40cf81f0a765800a8db0"), "7990000000000000000000000"},             /// New 7990000 Halo, Old 10000 Halo
		{common.HexToAddress("0x1e991f2c708e53219bdee9509edc2e0ac7ed7987"), "86691500000000000000000"},               /// New 86691.5 Halo, Old 108.5 Halo
		{common.HexToAddress("0x21a7a65a4f95d38e2532e8a757485d1db002239d"), "399500000000000000000000"},              /// New 399500 Halo, Old 500 Halo
		{common.HexToAddress("0x27f2876cba5d4304aac30bb2aa4a24934abfd73d"), "798201000000000000000000"},              /// New 798201 Halo, Old 999 Halo
		{common.HexToAddress("0x2dd8ee0f4d060d47c95f03ca488f30ac24db1e45"), "3995000000000000000000000"},             /// New 3995000 Halo, Old 5000 Halo
		{common.HexToAddress("0x32af90c37b9489487a6cc5a943c0039a50dfcded"), "4794000000000000000000000"},             /// New 4794000 Halo, Old 6000 Halo
		{common.HexToAddress("0x383fabb6a199254975138846fb0df217f3937a1e"), "7909301000000000000000000"},             /// New 7909301 Halo, Old 9899 Halo
		{common.HexToAddress("0x3aa470ccfb5fba3ca23a0a22cebbee79c7866377"), "719100000000000000000000"},              /// New 719100 Halo, Old 900 Halo
		{common.HexToAddress("0x3c773cd2e8b581c7832e1484a77424f3e4d26b08"), "3995000000000000000000000"},             /// New 3995000 Halo, Old 5000 Halo
		{common.HexToAddress("0x3d0863e07c6a592d425caeecedea921a53c7e825"), "799000000000000000000"},                 /// New 799 Halo, Old 1 Halo
		{common.HexToAddress("0x41e2919b67c406c9d76b102881f937d185c9824d"), "3595500000000000000000000"},             /// New 3595500 Halo, Old 4500 Halo
		{common.HexToAddress("0x421c682b6e6164aa1de7fb588d89f5f44e451e3e"), "719100000000000000000000"},              /// New 719100 Halo, Old 900 Halo
		{common.HexToAddress("0x442b6ec73476cb132a9d6d6935493449bb775a5f"), "1440597000000000000"},                   /// New 1.440597 Halo, Old 0.001803 Halo
		{common.HexToAddress("0x444b6915e71a9ec9b2ed7298bb0083d90d24ec85"), "39151000000000000000000000000"},         /// New 39151000000 Halo, Old 49000000 Halo
		{common.HexToAddress("0x48c469abee81c6e28e4fa92ab59514eea9c6903d"), "98641967220000000000"},                  /// New 98.64196722 Halo, Old 0.12345678 Halo
		{common.HexToAddress("0x49c208f64ef73760a1cee1964a0f4f98ab334f8c"), "384239100000000000000000000"},           /// New 384239100 Halo, Old 480900 Halo
		{common.HexToAddress("0x4b08c1e3c96caa920016ffcaa4ad73df1a3474de"), "3195999494077712560240000000000"},       /// New 3195999494077.71256024 Halo, Old 3999999366.80564776 Halo
		{common.HexToAddress("0x4b19ee2d9edfd680b4ffeeaf34252edfbf213148"), "7990000000000000000000"},                /// New 7990 Halo, Old 10 Halo
		{common.HexToAddress("0x4e13cbd86cb66985b13a29a2394a171c41fc1d4d"), "7990000000000000000000"},                /// New 7990 Halo, Old 10 Halo
		{common.HexToAddress("0x4fbe7a074bf7ec4235b58d9ea307a5859f942c97"), "3196000000000000000000000"},             /// New 3196000 Halo, Old 4000 Halo
		{common.HexToAddress("0x5574e56c5946188b7719bcf9606fa41ace2d7330"), "79388640000000000000000000"},            /// New 79388640 Halo, Old 99360 Halo
		{common.HexToAddress("0x593855cac489d6f164431f8bbadf24d2c4594f72"), "799000000000000000000000"},              /// New 799000 Halo, Old 1000 Halo
		{common.HexToAddress("0x60cf9a1aafff050f1714b8d703e8eab56bd58384"), "39950000000000000000000000000"},         /// New 39950000000 Halo, Old 50000000 Halo
		{common.HexToAddress("0x61e9991020c29772862c04987be04a2885524a5e"), "710638590000000000000000000000"},        /// New 710638590000 Halo, Old 889410000 Halo
		{common.HexToAddress("0x661bacdf5aadcb07a1314ab2e067b98934f7a6eb"), "397532462500000000000000000"},           /// New 397532462.5 Halo, Old 497537.5 Halo
		{common.HexToAddress("0x67059715c4a6245c6d47c1d20dd795694ead0745"), "24009950000000000000000000"},            /// New 24009950 Halo, Old 30050 Halo
		{common.HexToAddress("0x6786836317e095fcd976af9c8918dc36e5e2f04a"), "161645290500000000000000000"},           /// New 161645290.5 Halo, Old 202309.5 Halo
		{common.HexToAddress("0x67b7301878bc38559e118e9958d1bb8b4eb6afb2"), "7190192115404942975500079899840200000"}, /// New 7190192115404942975.5000798998402 Halo, Old 8998988880356624.5000000999998 Halo
		{common.HexToAddress("0x6bbe54361de26f486ac44aa65076081a4b854c6f"), "7990077503000000000000000000"},          /// New 7990077503 Halo, Old 10000097 Halo
		{common.HexToAddress("0x6c5927ddc642913351a06ded29df10b58bd6c715"), "7964307098998853500000000000"},          /// New 7964307098.9988535 Halo, Old 9967843.6783465 Halo
		{common.HexToAddress("0x7009794cb3694f27af33a5b5ef9fe1b636472ec4"), "10626700000000000"},                     /// New 0.0106267 Halo, Old 0.0000133 Halo
		{common.HexToAddress("0x79026249c10e8e1383694f38025bde7c2b48b99e"), "87755521109000000000000000"},            /// New 87755521.109 Halo, Old 109831.691 Halo
		{common.HexToAddress("0x79cab0202f4956441d54ade87b9c14ad18756464"), "39950000000000000000000000"},            /// New 39950000 Halo, Old 50000 Halo
		{common.HexToAddress("0x7af63f548154f92330edfafcaa16bf16236f521e"), "806965933318603000000000000000000"},     /// New 806965933318603 Halo, Old 1009969878997 Halo
		{common.HexToAddress("0x7be840c83eb70d560819c486029bbbd682f44a37"), "15181000000000000000000"},               /// New 15181 Halo, Old 19 Halo
		{common.HexToAddress("0x7ca09c776bd4be51cb466600c85e77b0a9d37b03"), "453407731000000000000000000"},           /// New 453407731 Halo, Old 567469 Halo
		{common.HexToAddress("0x7f3523dc13ffb78acd0ed853727c5967f998c593"), "7989995146075000000000000"},             /// New 7989995.146075 Halo, Old 9999.993925 Halo
		{common.HexToAddress("0x897930f4f33bc53694254c8be0ecf8d8718162e6"), "79100112311000000000000000"},            /// New 79100112.311 Halo, Old 98998.889 Halo
		{common.HexToAddress("0x90171a5e9f1782a85b5aac71bdb992d61d1de089"), "799000000000000000000"},                 /// New 799 Halo, Old 1 Halo
		{common.HexToAddress("0x91e0be4f8a30d9a1eab2495a25e0f00c1eb42897"), "799000000000000000000000000"},           /// New 799000000 Halo, Old 1000000 Halo
		{common.HexToAddress("0x928e5528aaa3c19640836c802c2743a0346acc56"), "7190993093612803531555406211270907560"}, /// New 7190993093612803531.55540621127090756 Halo, Old 8999991356211268.50006934444464444 Halo
		{common.HexToAddress("0x9734e83431b9a51958be493b07898fe1e1ca59c6"), "7990000000000000000000000"},             /// New 7990000 Halo, Old 10000 Halo
		{common.HexToAddress("0x981f7384d6f1cd570c57054f93aaa6cfa6f9200e"), "39950000000000000000000"},               /// New 39950 Halo, Old 50 Halo
		{common.HexToAddress("0x9bb163f86f10ca63ff7d09f18b8a8b8da4be0084"), "480832607000000000000000000"},           /// New 480832607 Halo, Old 601793 Halo
		{common.HexToAddress("0x9c5630cc365c1b6956356772f5a908a091fbcb7d"), "657760770000000000000000000"},           /// New 657760770 Halo, Old 823230 Halo
		{common.HexToAddress("0xa0861af7a2443dcbd3d5d4ac3d392b7d6c0c37bf"), "7590500000000000000000000"},             /// New 7590500 Halo, Old 9500 Halo
		{common.HexToAddress("0xa160ddb5f7a6dab43dd10e0e093f5a872db95a0e"), "968787500000000000"},                    /// New 0.9687875 Halo, Old 0.0012125 Halo
		{common.HexToAddress("0xa8ccd69bf9693e3bf44f0fda0491564ee41d09e5"), "505922287439760000000000"},              /// New 505922.28743976 Halo, Old 633.19435224 Halo
		{common.HexToAddress("0xa9da7b0279a00b4e11b8a7b4c72959e1eea496ca"), "799000000000000000000000"},              /// New 799000 Halo, Old 1000 Halo
		{common.HexToAddress("0xaaebff0e2d5583a358702e4863f89bc576673bd7"), "110806206890000000000000"},              /// New 110806.20689 Halo, Old 138.68111 Halo
		{common.HexToAddress("0xad31a2fa7d7e59ff65a8ad3b9f08c940ed4f5557"), "159735281000000000000000000000"},        /// New 159735281000 Halo, Old 199919000 Halo
		{common.HexToAddress("0xaddccf4b37a43503fc2133e21eca04ccab5e4818"), "796416593300000000000000000000"},        /// New 796416593300 Halo, Old 996766700 Halo
		{common.HexToAddress("0xae432f83fd83fbc79845a63bdc1ec382754c6251"), "2397000000000000000000"},                /// New 2397 Halo, Old 3 Halo
		{common.HexToAddress("0xb5e1be462a1d279c326a01d6ba218c491df7f7ea"), "7190201000000000000000000"},             /// New 7190201 Halo, Old 8999 Halo
		{common.HexToAddress("0xb5efab81110dbc0ccf4628b2e41c48bfba581c5e"), "7191000000000000000000000"},             /// New 7191000 Halo, Old 9000 Halo
		{common.HexToAddress("0xbcd47dbcf3ba3859b674703ac2cc2ffd63173b40"), "759048402000000000000000000"},           /// New 759048402 Halo, Old 949998 Halo
		{common.HexToAddress("0xc264e5040c1f94cf9753b3dcc9b0f2a70000a29d"), "799000000000000000000"},                 /// New 799 Halo, Old 1 Halo
		{common.HexToAddress("0xc58d8fbc5775375671f5c7c2ba0d5ddd0018ce52"), "7990000000000000000000000"},             /// New 7990000 Halo, Old 10000 Halo
		{common.HexToAddress("0xc8016c038684607f6815abda9fd0d4d920cba11e"), "2841244000000000000000000"},             /// New 2841244 Halo, Old 3556 Halo
		{common.HexToAddress("0xc89065e6fffce7766bda032387ca0fb44183d3e6"), "78917068828129592025185302440"},         /// New 78917068828.12959202518530244 Halo, Old 98769798.28301575973114556 Halo
		{common.HexToAddress("0xd104ac05e7ccca9677e0714f7c88d3608a9b4f04"), "958800000000000"},                       /// New 0.0009588 Halo, Old 0.0000012 Halo
		{common.HexToAddress("0xe11fffd526b2afa0fbd0f89ed854e19146af9172"), "31960000000000000000000"},               /// New 31960 Halo, Old 40 Halo
		{common.HexToAddress("0xe29a4120dce8119135de3c7b0cd12bb81e23a156"), "798663931012000000000000000000"},        /// New 798663931012 Halo, Old 999579388 Halo
		{common.HexToAddress("0xe5f09c9a586ed646e0d418cb878bd104b1b9ec35"), "901527680000000000000000"},              /// New 901527.68 Halo, Old 1128.32 Halo
		{common.HexToAddress("0xe6fc4ab2c96cdfc3f25bd337a3ce8a696f21eda6"), "3196000000000000000000000"},             /// New 3196000 Halo, Old 4000 Halo
		{common.HexToAddress("0xe86df390af83fe236cdd073049a81ef11772705a"), "154938085000000000000000000"},           /// New 154938085 Halo, Old 193915 Halo
		{common.HexToAddress("0xe9e8b2df5fbe577c0d9c30fb1dc0936bbe90b6eb"), "76783900000000000000000000"},            /// New 76783900 Halo, Old 96100 Halo
		{common.HexToAddress("0xea39f1fd841b1d6c0d273be54af9fbb7855b92bd"), "355394401000000000000000000"},           /// New 355394401 Halo, Old 444799 Halo
		{common.HexToAddress("0xeaa3808ad902867b234826270a1092ab915ae6c5"), "9588000000000000000000"},                /// New 9588 Halo, Old 12 Halo
		{common.HexToAddress("0xec502eb6ffb14fc59c8b5a491bcab2558a398b7e"), "1789760000000000"},                      /// New 0.00178976 Halo, Old 0.00000224 Halo
		{common.HexToAddress("0xecaa05f7694b751029c49978b60369592e3c87c6"), "6987574360300000000000000000"},          /// New 6987574360.3 Halo, Old 8745399.7 Halo
		{common.HexToAddress("0xef0fd7f817f42863051cee9ccd34e227acc1d5bb"), "79900000000000000000000"},               /// New 79900 Halo, Old 100 Halo
		{common.HexToAddress("0xf4e2d5eba3cf7887ac2cd4cdb61251d53accec18"), "399499999999999201000"},                 /// New 399.499999999999201 Halo, Old 0.499999999999999 Halo
		{common.HexToAddress("0xfd22542017de3ed588b34c72da35950657ca6341"), "7990000000000000000000000"},             /// New 7990000 Halo, Old 10000 Halo
		{common.HexToAddress("0xfe733b82edd65d6e01c8d752b6f3943d9f3eff1e"), "71910000000000000000000"},               /// New 71910 Halo, Old 90 Halo
		{common.HexToAddress("0x00c2d83978a90d7ce47c140b8095830c4d7272dc"), "607240000000000"},                       /// New 0.00060724 Halo, Old 0.00000076 Halo
		{common.HexToAddress("0x04798aae1acf00aaa739630b69d2cc49cb78db77"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x06c06b33791908c3836614382190d0e42bb9f00d"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x07ffddd9957f4bf09a6b7065a584bcac2c572d80"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x0e6fd387cd5d3b5f55d367133da87fc3d4e0ea0b"), "399500000000000000000000"},              /// New 399500 Halo, Old 500 Halo
		{common.HexToAddress("0x0f71ca6fd33d6698ad090f0d42a7c2e7f79e84cc"), "6392000000001278400000"},                /// New 6392.0000000012784 Halo, Old 8.0000000000016 Halo
		{common.HexToAddress("0x11e0fd692f3a483ca8e2027305fd867f767eb7f9"), "159800000000000000000000000"},           /// New 159800000 Halo, Old 200000 Halo
		{common.HexToAddress("0x132cabda25193fb0388d18033e2a0f9a91102b4d"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x1405454131b4991206489e85abc9ad83000cff7e"), "399500000000000000000000"},              /// New 399500 Halo, Old 500 Halo
		{common.HexToAddress("0x1650349a598297347952680a45d1acab01a630c5"), "799000000000000"},                       /// New 0.000799 Halo, Old 0.000001 Halo
		{common.HexToAddress("0x1799c1e742b0b840a421d606c50b3493fb701103"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x17f23f40f9d8b70c484b6b0175b41bf43deba485"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x189e26489cd9b590862e80acd048be1e301b6977"), "985966000000000000000000"},              /// New 985966 Halo, Old 1234 Halo
		{common.HexToAddress("0x1a085d3bed57e32bc476b7de465d8b6e35471ef9"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x1abcbd50b49ddc8faf1806ceb5398de85681d1b0"), "1598000000000000000000"},                /// New 1598 Halo, Old 2 Halo
		{common.HexToAddress("0x1bb339729962652cbc8922d0e954affa43f6bb43"), "1598000000000000000000"},                /// New 1598 Halo, Old 2 Halo
		{common.HexToAddress("0x1bf138cd1c4b8c8a4641984f0247f5d94e38fe0b"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x1ff3faaae88d35236a6389a1a2b770b1e3630b92"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x2176af3597ce5ac7a185bd2c51a4476729faf090"), "4794000000000000000000"},                /// New 4794 Halo, Old 6 Halo
		{common.HexToAddress("0x24c85a7a9407953ce043436250f2359ab9c7ead6"), "3995000000000000000000000"},             /// New 3995000 Halo, Old 5000 Halo
		{common.HexToAddress("0x27d11ad207ed227b80dd2b70abe846553994eeb5"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x28e8b23e5b0b9971b394e6a123781a8bfb92d2a2"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x2e2f85c75ee2c02bf78a9e8a168e04a8f9734dc1"), "799000000000000000000000"},              /// New 799000 Halo, Old 1000 Halo
		{common.HexToAddress("0x307de5d20189c049c012719ef30a0a7001842539"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x3412851d93610e9290d13afc4bc4117adc36572a"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x361e793bd53d231b16621dbda389c0e285d83c59"), "11985000000000000000000"},               /// New 11985 Halo, Old 15 Halo
		{common.HexToAddress("0x38f6e119e999e99ca5eadca1f5dbd6101f3d68f4"), "799000000000000000000"},                 /// New 799 Halo, Old 1 Halo
		{common.HexToAddress("0x3b39df9f0489021e42cc3c9e6713f6bbe421eecf"), "4794000000000000"},                      /// New 0.004794 Halo, Old 0.000006 Halo
		{common.HexToAddress("0x3b99541a89dc3bf5c30c4ad68fba7b96c4ee4275"), "7990000000000000000000000"},             /// New 7990000 Halo, Old 10000 Halo
		{common.HexToAddress("0x3cad65ceebf6218eea45c3ad24e8ca5eb0181ec3"), "2397000000000000"},                      /// New 0.002397 Halo, Old 0.000003 Halo
		{common.HexToAddress("0x3ec1fc1fdb85bd3e131067b6bb914b047aa76a1d"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x418debe1c89bdf68cc636217ae0f78306e29e113"), "838950000000000000000000000"},           /// New 838950000 Halo, Old 1050000 Halo
		{common.HexToAddress("0x423c17fafc6dcc08a0fa3bf715e161d747dcd504"), "799000000000000000000000000"},           /// New 799000000 Halo, Old 1000000 Halo
		{common.HexToAddress("0x4274b1d6485c23b60dd3063ef1d175d650d94a66"), "1598000000000000000000"},                /// New 1598 Halo, Old 2 Halo
		{common.HexToAddress("0x436e5eac09325199d3023234ac37e9b66cbfdd29"), "3995000000000000000000"},                /// New 3995 Halo, Old 5 Halo
		{common.HexToAddress("0x4827eef69cc194f6b64a1ae8e37160ea7361d169"), "80699000000000000000000"},               /// New 80699 Halo, Old 101 Halo
		{common.HexToAddress("0x4c78b134470554a07db8777c3392dd784cf6dc22"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x54ed3a5f70a3ead37680821404a52f8d10381047"), "799000000000000000000000"},              /// New 799000 Halo, Old 1000 Halo
		{common.HexToAddress("0x590a473bde0e790574e8c7416f6c0b029855abf5"), "1598000000000000"},                      /// New 0.001598 Halo, Old 0.000002 Halo
		{common.HexToAddress("0x59dd84399ec0ef160e561dbe8185b1c823811e4f"), "799000000000000"},                       /// New 0.000799 Halo, Old 0.000001 Halo
		{common.HexToAddress("0x5a0f4fb7c0dc3f3c22908af12f2dff50ecc82a3a"), "43146000000000000000000"},               /// New 43146 Halo, Old 54 Halo
		{common.HexToAddress("0x5a2c9fcb76cb6e2459689c76ac3da652e1515771"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x5a4fa35f34dcf6d50554138951ee9d6d9dcc6318"), "39950000000000000000"},                  /// New 39.95 Halo, Old 0.05 Halo
		{common.HexToAddress("0x609aad4174514ba68397e2d414c78c10d362e395"), "1598000000000000000000"},                /// New 1598 Halo, Old 2 Halo
		{common.HexToAddress("0x613c5ec41816ac81698ac1b8a8ea0209dae597c5"), "399500000000000000000000"},              /// New 399500 Halo, Old 500 Halo
		{common.HexToAddress("0x627761724d0421f4735f9fb9104e9dfa08f6c1e9"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x633062bf6cdce3b62e75772419c3b0bf895d73ca"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x64f172b0a66ed2c3b10b97c1f2c112ead65375c4"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x656d9e9586efb9c59e1cd3e53775719771fbad40"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x683fd0c2f9f4f8bbe0faf62799db95fa67df6f49"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x6a9360a308056ff75697dc2cac8e7fefd62ccc80"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x6cc9eec29d14debf91fdbf82c1b2ea0238ef898a"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x6d3095660d2aa4692d8004f47c1303fe264cf8f5"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x6f6bb33e58560aa77f11deb7fca398055dabbea2"), "95880000000000000000000000"},            /// New 95880000 Halo, Old 120000 Halo
		{common.HexToAddress("0x6f837f15cb5bd065a24b5c0e252603c5b24818db"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x70b315e75067468ff3e5252f3185bb5ccd3cb606"), "1598000000000000000000"},                /// New 1598 Halo, Old 2 Halo
		{common.HexToAddress("0x764a2d04a8fce5e205a2692672563b8ee87aac96"), "799000000000000000000000"},              /// New 799000 Halo, Old 1000 Halo
		{common.HexToAddress("0x7843f951032493c3c61392dceba6ec2529de5bf8"), "561609819653521532079437243"},           /// New 561609819.653521532079437243 Halo, Old 702890.888177123319248357 Halo
		{common.HexToAddress("0x7c6cdc420245e860fc04399741a95c9225b542ac"), "1598000000000000000000"},                /// New 1598 Halo, Old 2 Halo
		{common.HexToAddress("0x7f1bda97ebcda275b7b734ea20d95e7a00edc1c5"), "799000000000000"},                       /// New 0.000799 Halo, Old 0.000001 Halo
		{common.HexToAddress("0x83165cbb0bc4f1f9836677ccfde2a6d38355dfbf"), "119850000000000000000000"},              /// New 119850 Halo, Old 150 Halo
		{common.HexToAddress("0x83d191a52c6a7b9d6add9fede8b13b87f6f2fae8"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x880b01a259e5557abc50372348a9c5fa002db568"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x89af887aabdf3e26241c696175f7db72614fc9cb"), "3995000000000000000000"},                /// New 3995 Halo, Old 5 Halo
		{common.HexToAddress("0x8b962c4f6b22ec9b97c15927a5858f82211aa5a0"), "303620000000000"},                       /// New 0.00030362 Halo, Old 0.00000038 Halo
		{common.HexToAddress("0x8c5ab8afcfedcde5c3c7c58d4bacb88352a2893b"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x8d0988ce7d528efff1ce84013182706ec267eed4"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x8d4609ab9bba0a2e6998408acc145edc3668498e"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x8d4ad8fa781d43898cbffcf1790e49922bdc8b8a"), "399500000000000000000000"},              /// New 399500 Halo, Old 500 Halo
		{common.HexToAddress("0x8d79b150e8e1f2156e623bda94a005c22edad65a"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x8eed335a850576afbd1219a26c1e12f3326b4499"), "14382000000000000"},                     /// New 0.014382 Halo, Old 0.000018 Halo
		{common.HexToAddress("0x90743eb362a64c19ec09eca4ed8fced68c35e81a"), "399500000000000000000000"},              /// New 399500 Halo, Old 500 Halo
		{common.HexToAddress("0x9104ade6d31058d0f23d75f5239c0bcbce498df3"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x919832842b1ac603f81d0d164968ff7b4758b406"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x93e6ba2538cb88ab6e7828464f017175a248cdf0"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0x961b99b59a01fe76c9b2d38898ae867fc4ea2a7c"), "799000000000000"},                       /// New 0.000799 Halo, Old 0.000001 Halo
		{common.HexToAddress("0x9b2e19c550e7e27b684135f485fe0ff4d047acab"), "399500000000000000000000"},              /// New 399500 Halo, Old 500 Halo
		{common.HexToAddress("0x9b6ccaa138a4c34b884deeb6b6c82ef645e783eb"), "399500000000000000000000"},              /// New 399500 Halo, Old 500 Halo
		{common.HexToAddress("0x9ca97000790c2c92b08ac45bbdf75860b6350af2"), "878900000000000000000000000"},           /// New 878900000 Halo, Old 1100000 Halo
		{common.HexToAddress("0x9fa03fc6587d3823de4f77de81618097b0b8db1a"), "127840000000000000000000"},              /// New 127840 Halo, Old 160 Halo
		{common.HexToAddress("0xa09e12fbdda18422a963b5355ce7178509b7d308"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xa2ef713e708865a7b940c7f57c93453c618acfa0"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xa413ad4ba22d5785c026327fdf6cc03438eaa908"), "1598000000000000000000"},                /// New 1598 Halo, Old 2 Halo
		{common.HexToAddress("0xa41fc486303f58b5775f38a7ebc3af341bd3fff4"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xa5bbc0b1288c717caf8350a760c9d147d3974f90"), "7990000000000000000000000"},             /// New 7990000 Halo, Old 10000 Halo
		{common.HexToAddress("0xa7147322cc5e7743a51533383bdcf176beb2958e"), "39950000987414754790000"},               /// New 39950.00098741475479 Halo, Old 50.00000123581321 Halo
		{common.HexToAddress("0xaafb0cad320d9a689144ed855e16d77a1de71c47"), "7990000000000000000000000"},             /// New 7990000 Halo, Old 10000 Halo
		{common.HexToAddress("0xabb427aecc0553f1989ea15250a18468cf7660b1"), "798201000000000000000000"},              /// New 798201 Halo, Old 999 Halo
		{common.HexToAddress("0xaff4b62628169944233f39e6e98bbe4577e7e1c8"), "98641967220000000000"},                  /// New 98.64196722 Halo, Old 0.12345678 Halo
		{common.HexToAddress("0xb4fd3a1c8d16b05983910831139fde160e0efa10"), "1598000000000000000000"},                /// New 1598 Halo, Old 2 Halo
		{common.HexToAddress("0xb62e10c2a1687ce9d415b7bd8f28626a7bfbb8c1"), "799799000000000000000000"},              /// New 799799 Halo, Old 1001 Halo
		{common.HexToAddress("0xb63baf580f4ac8055772db6f11a877a526fc24ef"), "7990000000000000000000000"},             /// New 7990000 Halo, Old 10000 Halo
		{common.HexToAddress("0xb681f3b2dd25fdb393df7ba2824eaff5045ee9e3"), "799000000000000000000000"},              /// New 799000 Halo, Old 1000 Halo
		{common.HexToAddress("0xb8ff4ee87dc98bb6655e10eb01be4fa3b3f6ccb0"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xb90bb6dbc5b519f7119eaa0f46084cac468d6c28"), "6392000000007670400000"},                /// New 6392.0000000076704 Halo, Old 8.0000000000096 Halo
		{common.HexToAddress("0xbd840ce9b4fdedb2e7026d77ba881e0975992e38"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xbdffff801ef19729303ca8d1ff617991fad08a73"), "79900000000000000000000000"},            /// New 79900000 Halo, Old 100000 Halo
		{common.HexToAddress("0xbea7a50ffd19adcab016712c3ca7bae176eff2fd"), "2397000000000000000000"},                /// New 2397 Halo, Old 3 Halo
		{common.HexToAddress("0xc2f9526a1629cd3b1ee322c2f85e07e2af65719e"), "799000000000000000000000"},              /// New 799000 Halo, Old 1000 Halo
		{common.HexToAddress("0xc75458ccfcd3aba31a02a8d05495aabb0d68c6ab"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xc7bb232c36fd346e2156867de122f11a14658b93"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xc9b7cc37eca0f2cb271451e9f9b14c817d5e3233"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xc9fd8e98fc65769ecf3b0286912a2e276c9179d3"), "3995000000000000000000000000"},          /// New 3995000000 Halo, Old 5000000 Halo
		{common.HexToAddress("0xca37dbcbc79d82a50648419cf114b05fdef87844"), "84562423155470393618926287150"},         /// New 84562423155.47039361892628715 Halo, Old 105835323.09821075546799285 Halo
		{common.HexToAddress("0xcc72446ae75043e6fe3a2d277e35a075c98ff687"), "1598000000000000000000"},                /// New 1598 Halo, Old 2 Halo
		{common.HexToAddress("0xd27d6a1dd27201c19001fec323e42dc0720b6f4f"), "799887689000000000000000"},              /// New 799887.689 Halo, Old 1001.111 Halo
		{common.HexToAddress("0xd2f216e89e7ecdf186363934e81cb675160e3783"), "3995000000000000000000"},                /// New 3995 Halo, Old 5 Halo
		{common.HexToAddress("0xd4db34742498d404adaa255e344b48e0ffae1a41"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xd51bc95094fe13a1871b56273de342adc19df692"), "2397799000000000000"},                   /// New 2.397799 Halo, Old 0.003001 Halo
		{common.HexToAddress("0xde0c7541f019d056cc220c3eaac6be369420a112"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xde70f70056153276821f1f4d1019a614fad8878e"), "8789000000000000000000000"},             /// New 8789000 Halo, Old 11000 Halo
		{common.HexToAddress("0xdfd285cd5e3fadc4de5c1d5a9f4c735c20ddd119"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xdfeeec9581787690461d03090b7be3a78ac5e042"), "5593000000000000"},                      /// New 0.005593 Halo, Old 0.000007 Halo
		{common.HexToAddress("0xe1b07e030037a393fae43de20f496fbad362576b"), "105048657824858081609002108561"},        /// New 105048657824.858081609002108561 Halo, Old 131475166.238871191000002639 Halo
		{common.HexToAddress("0xe7e8ebbe46b5ed6d5d0aaba9a9f3cfa1c3784b73"), "79900000000000000000000000"},            /// New 79900000 Halo, Old 100000 Halo
		{common.HexToAddress("0xe927c9361c4ce0a5115b1f188ca41b303b50380b"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xe97d51c1aece36f9188f7d92fa63b0718ad1721f"), "799000000000000000000000000000"},        /// New 799000000000 Halo, Old 1000000000 Halo
		{common.HexToAddress("0xe9905cd30c7dfdccd65d402f3d32b2d6dacc15a8"), "7990000000000000000000000"},             /// New 7990000 Halo, Old 10000 Halo
		{common.HexToAddress("0xea746d971364adeead842e20ae3a25cce47e1e1b"), "7590500000000000000000000"},             /// New 7590500 Halo, Old 9500 Halo
		{common.HexToAddress("0xeb747b636fa2d69c29032f214da823b6cfdc92ed"), "7990000000000000000000000"},             /// New 7990000 Halo, Old 10000 Halo
		{common.HexToAddress("0xebb0bb4e2c583acce839e7a32a5e49bdcd5c0b1f"), "399500000000000000000000"},              /// New 399500 Halo, Old 500 Halo
		{common.HexToAddress("0xec2641e1a1eb681294a97e8f2578772ecd29effb"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xed2eb93625f7c3a7d3232eaabdf615e478c99847"), "7990000000000000000000000"},             /// New 7990000 Halo, Old 10000 Halo
		{common.HexToAddress("0xed38c32b3dfcdc550bb83c6aad6ff0a3f81501a9"), "399500000000000000000000000"},           /// New 399500000 Halo, Old 500000 Halo
		{common.HexToAddress("0xeea6a6b2df4176dc851d6393cdded77684c324d8"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xf13ed5f7d1e35a2695a3b609ab81d33f4ca3ed6b"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xf15a0f1343e1aa96a6f576287ac7b05a8e8e3180"), "7990000000000000000000000"},             /// New 7990000 Halo, Old 10000 Halo
		{common.HexToAddress("0xf20cfca236a56f0c707f92c4f03b786b0d78edc9"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xf2ca43f48981e3b1a84e1ce79c25a6296449602a"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xf4a610f13bdbd6cc1acc17b692fb584836e79027"), "443445000000000000000000"},              /// New 443445 Halo, Old 555 Halo
		{common.HexToAddress("0xf5431d96b6be6a5233a5b654b6de83e058d5697e"), "1598000000000000000000"},                /// New 1598 Halo, Old 2 Halo
		{common.HexToAddress("0xf57548407b3a4889749c7fce2e508b02156d3465"), "7990000000000000000000000"},             /// New 7990000 Halo, Old 10000 Halo
		{common.HexToAddress("0xf89dc950a93503279930c4047682a743833e2c3b"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xfb3e598cd74bf511baba6e31ca24eb0b9714770c"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xfb5702df871a63ed54f949d54270637f5472433a"), "1481098813892146500000000000"},          /// New 1481098813.8921465 Halo, Old 1853690.6306535 Halo
		{common.HexToAddress("0xfba186bfcf9360c3d2315edbd2a0438a858fdcdd"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xfd29674b0008bd5de491281dd3f38584636b2478"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xfd3cbcc486acfb19c5278f96d52831bc9fb6ebb4"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xfed3b709f1d03a03255ce8776911a1656235ef20"), "6392000000000000000000"},                /// New 6392 Halo, Old 8 Halo
		{common.HexToAddress("0xffb453ff49f6faf7e713b2fa5b636d9b4877440c"), "808133405325640863239992966046"},        /// New 808133405325.640863239992966046 Halo, Old 1011431045.463880930212757154 Halo
	}
}

// Account list to inflate for local development
func LocalCoinSplitAccount() []AccountInfo {
	return []AccountInfo{
		{common.HexToAddress("0x67b7301878bc38559e118e9958d1bb8b4eb6afb2"), "499000000000000000000"}, /// 199 Halo
		{common.HexToAddress("0x928e5528aaa3c19640836c802c2743a0346acc56"), "599000000000000000000"}, /// 299 Halo
	}
}
