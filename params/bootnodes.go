// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Silesiacoin Foundation Go Bootnodes
	"enode://4470ab64fa94cc854ba4f79e7b846cf3a5e2d960f13bb103c2c8878885ddfae422d96a8739929382ae2de1522a5bf9ec2825b77dab8d839e7a43500c9732882b@bootnode1.silesiacoin.com:30303", // bootnode-aws-1
	"enode://769c6f081e11c4bc4e4436c93c373f44e1226b8ee2fde086ac8eac38c748fd6dd4558f5d86cfba8f1012d34a352369c10e2879a534ade9da38cf10527186c799@bootnode2.silesiacoin.com:30303", // bootnode-gcp-1
	"enode://b5f7450bdd598c0048a46e7e345bdf426876d9ff37127976ef16080ed7641dd8279406bb82dccc0125afddc0f0c0ae87c6ed5c36fcd25f59a365ad1859b5abf8@bootnode3.silesiacoin.com:30303", // bootnode-local-1
	"enode://445e870d61c96bbd34211437222d09d97c4b2da99ed5f0832d0869a764b10377b13825d742b92f5de227b2517318aca51b779f5944ea5ce6289762df1c901f86@bootnode4.silesiacoin.org:30303", // bootnode-gcp-2
	"enode://0ed7807ffd693dffa0d93dddbec5e864365105324aed35d81b753eb47382d50edd1e3246709163834994a4bcd1f9b4bd09b35da2a187d7d0bf85e098bcbf8d19@bootnode5.silesiacoin.org:30303", // bootnode-local-2
	"enode://47106ba1e12d4506831f750165bd4a92fec3480dc3c4d592fd255c7a3a32c51fd8e788c79b9b7ca1ac142c77bd1c1d0d32fcaac4214ebc4a4ec6d2e2a5eadeed@bootnode6.silesiacoin.org:30303", // bootnode-aws-2
	//	"enode://5c444c192be39dab5a8fbaf15654f19aacc78500fb63f8e02deee896cd7bf33fce769665e4600358d66fba2273db86e971b9beb9bbc6e9ec4a6834162975d9ad@bootnode7.ssc:30303",             // bootnode-do-1
	//	"enode://bc154cee3a01fd31258df5927a4999e2441aad024667266d4e92aec828dc652b922db8f33f2fff75d674b3dfbdbc1c9de3d47c8223d8a63bd59edd2950a048e2@bootnode8.ssc:30303",             // bootnode-do-2
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Silesiacoin test network.
var TestnetBootnodes = []string{
	//"enode://d26ea46056e927a3de8ea0268ee472be654170f2de34b6fceb87f9474b3f2c68e1b0905ed555fc1b4485a3800279ecb904010dc38092413c76f291c8c0e1563f@bootnode.ssc:30303",             // Local terra implementation
	"enode://d26ea46056e927a3de8ea0268ee472be654170f2de34b6fceb87f9474b3f2c68e1b0905ed555fc1b4485a3800279ecb904010dc38092413c76f291c8c0e1563f@testnode.silesiacoin.com:30303",  // Cloud terra implementation
	"enode://d26ea46056e927a3de8ea0268ee472be654170f2de34b6fceb87f9474b3f2c68e1b0905ed555fc1b4485a3800279ecb904010dc38092413c76f291c8c0e1563f@testnode1.silesiacoin.com:30303", // Silesia test node
	//"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	//"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://c1f8b7c2ac4453271fa07d8e9ecf9a2e8285aa0bd0c07df0131f47153306b0736fd3db8924e7a9bf0bed6b1d8d4f87362a71b033dc7c64547728d953e43e59b2@52.64.155.147:30303",
	"enode://f4a9c6ee28586009fb5a96c8af13a58ed6d8315a9eee4772212c1d4d9cebe5a8b8a78ea4434f318726317d04a3f531a1ef0420cf9752605a562cfe858c46e263@213.186.16.82:30303",

	// Ethereum Foundation bootnode
	"enode://573b6607cd59f241e30e4c4943fd50e99e2b6f42f9bd5ca111659d309c06741247f4f1e93843ad3e8c8c18b6e2d94c161b7ef67479b3938780a97134b618b5ce@52.56.136.200:30303",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://06051a5573c81934c9554ef2898eb13b33a34b94cf36b202b69fde139ca17a85051979867720d4bdae4323d4943ddf9aeeb6643633aa656e0be843659795007a@35.177.226.168:30303",
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30304",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30306",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30307",
}
