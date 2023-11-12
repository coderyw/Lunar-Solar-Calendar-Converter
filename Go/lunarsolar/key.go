// Package lunarsolar
// @Author: yinwei
// @File: key
// @Version: 1.0.0
// @Date: 2023/11/12 22:10

package lunarsolar

// 生成方法在上层 generate.html
var (
	lunar_month_days = [...]int{1887, 0x1694, 0x16aa, 0x4ad5, 0xab6, 0xc4b7,
		0x4ae, 0xa56, 0xb52a, 0x1d2a, 0xd54, 0x75aa, 0x156a, 0x1096d, 0x95c,
		0x14ae, 0xaa4d, 0x1a4c, 0x1b2a, 0x8d55, 0xad4, 0x135a, 0x495d, 0x95c, 0xd49b, 0x149a, 0x1a4a, 0xbaa5, 0x16a8, 0x1ad4, 0x52da, 0x12b6, 0xe937, 0x92e, 0x1496, 0xb64b, 0xd4a, 0xda8, 0x95b5, 0x56c, 0x12ae, 0x492f, 0x92e, 0xcc96, 0x1a94, 0x1d4a, 0xada9, 0xb5a, 0x56c, 0x726e, 0x125c, 0xf92d, 0x192a, 0x1a94, 0xdb4a, 0x16aa, 0xad4, 0x955b, 0x4ba, 0x125a, 0x592b, 0x152a, 0xf695, 0xd94, 0x16aa, 0xaab5, 0x9b4, 0x14b6, 0x6a57, 0xa56, 0x1152a, 0x1d2a, 0xd54, 0xd5aa, 0x156a, 0x96c, 0x94ae, 0x14ae, 0xa4c, 0x7d26, 0x1b2a, 0xeb55, 0xad4, 0x12da, 0xa95d, 0x95a, 0x149a, 0x9a4d, 0x1a4a, 0x11aa5, 0x16a8, 0x16d4, 0xd2da, 0x12b6, 0x936, 0x9497, 0x1496, 0x1564b, 0xd4a, 0xda8, 0xd5b4, 0x156c, 0x12ae, 0xa92f, 0x92e, 0xc96, 0x6d4a, 0x1d4a, 0x10d65, 0xb58, 0x156c, 0xb26d, 0x125c, 0x192c, 0x9a95, 0x1a94, 0x1b4a, 0x4b55, 0xad4, 0xf55b, 0x4ba, 0x125a, 0xb92b, 0x152a, 0x1694, 0x96aa, 0x15aa, 0x12ab5, 0x974, 0x14b6, 0xca57, 0xa56, 0x1526, 0x8e95, 0xd54, 0x15aa, 0x49b5, 0x96c, 0xd4ae, 0x149c, 0x1a4c, 0xbd26, 0x1aa6, 0xb54, 0x6d6a, 0x12da, 0x1695d, 0x95a, 0x149a, 0xda4b, 0x1a4a, 0x1aa4, 0xbb54, 0x16b4, 0xada, 0x495b, 0x936, 0xf497, 0x1496, 0x154a, 0xb6a5, 0xda4, 0x15b4, 0x6ab6, 0x126e, 0x1092f, 0x92e, 0xc96, 0xcd4a, 0x1d4a, 0xd64, 0x956c, 0x155c, 0x125c, 0x792e, 0x192c, 0xfa95, 0x1a94, 0x1b4a, 0xab55, 0xad4, 0x14da, 0x8a5d, 0xa5a, 0x1152b, 0x152a, 0x1694, 0xd6aa, 0x15aa, 0xab4, 0x94ba, 0x14b6, 0xa56, 0x7527, 0xd26, 0xee53, 0xd54, 0x15aa, 0xa9b5, 0x96c, 0x14ae, 0x8a4e, 0x1a4c, 0x11d26, 0x1aa4, 0x1b54, 0xcd6a, 0xada, 0x95c, 0x949d, 0x149a, 0x1a2a, 0x5b25, 0x1aa4, 0xfb52, 0x16b4, 0xaba, 0xa95b, 0x936, 0x1496, 0x9a4b, 0x154a, 0x136a5, 0xda4, 0x15ac, 0xcab6, 0x126e, 0x92e, 0x8c97, 0xa96, 0xd4a, 0x6da5, 0xd54, 0xf56a, 0x155a, 0xa5c, 0xb92e, 0x152c, 0x1a94, 0x9d4a, 0x1b2a, 0x16b55, 0xad4, 0x14da, 0xca5d, 0xa5a, 0x151a, 0xba95, 0x1654, 0x16aa, 0x4ad5, 0xab4, 0xf4ba, 0x14b6, 0xa56, 0xb517, 0xd16, 0xe52, 0x96aa, 0xd6a, 0x165b5, 0x96c, 0x14ae, 0xca2e, 0x1a2c, 0x1d16, 0xad52, 0x1b52, 0xb6a, 0x656d, 0x55c, 0xf45d, 0x145a, 0x1a2a, 0xda95, 0x16a4, 0x1ad2, 0x8b5a, 0xab6, 0x1455b, 0x8b6, 0x1456, 0xd52b, 0x152a, 0x1694, 0xb6aa, 0x15aa, 0xab6, 0x64b7, 0x8ae, 0xec57, 0xa56, 0xd2a, 0xcd95, 0xb54, 0x156a, 0x8a6d, 0x95c, 0x14ae, 0x4a56, 0x1a54, 0xdd2a, 0x1aaa, 0xb54, 0xb56a, 0x14da, 0x95c, 0x74ab, 0x149a, 0xfa4b, 0x1652, 0x16aa, 0xcad5, 0x5b4, 0x12ba, 0x895b, 0x936, 0x13497, 0xc96, 0xd52, 0xd6a9, 0xd6a, 0x56c, 0x92b6, 0x126e, 0x92e, 0x6c96, 0x1c94, 0xfd4a, 0x1b52, 0xb6a, 0xa56d, 0x55c, 0x125c, 0x992d, 0x192a, 0x13a95, 0x1694, 0x16d2, 0xeada, 0xab6, 0x4ba, 0xb25b, 0x1256, 0x152a, 0x7a95, 0x1694, 0x116aa, 0x15aa, 0xab6, 0xa4b7, 0x4ae, 0xa56, 0x952b, 0xd2a, 0x16d95, 0xb54, 0x156a, 0xc96d, 0x95c, 0x14ae, 0xaa4d, 0x1a4c, 0x1d2a, 0x6d55, 0xb54, 0xf55a, 0x12ba, 0x95c, 0xd49b, 0x149a, 0x1a4a, 0xbb25, 0x16a8, 0x1ad4, 0x32da, 0x12b6, 0xe957, 0x936, 0x1496, 0xb64b, 0xd4a, 0x15a8, 0x76b5, 0x56c, 0x112b6, 0x126e, 0x92e, 0xcc96, 0x1c94, 0x1d4a, 0x8da9, 0xb5a, 0x56c, 0x52ae, 0x125c, 0xd92d, 0x192a, 0x1a94, 0xbb4a, 0x16ca, 0xad4, 0x755b, 0x4ba, 0xf25b, 0x1256, 0x152a, 0xd695, 0xe94, 0x16aa, 0x8ad5, 0x9b4, 0x14b6, 0x4a57, 0xa56, 0xd52a, 0x1d2a, 0xd54, 0xb5aa, 0x156a, 0x96c, 0x74ae, 0x14ae, 0xea4d, 0x1a4c, 0x1b2a, 0xcd55, 0xad4, 0x135a, 0x895d, 0x95a, 0x1549b, 0x149a, 0x1a4a, 0xfaa5, 0x16a8, 0x16d4, 0xb2da, 0x12b6, 0x936, 0x7497, 0x1496, 0x1164b, 0xd4a, 0xda8, 0xd6b4, 0x156c, 0x12ae, 0x892f, 0x92e, 0xc96, 0x6d4a, 0x1d4a, 0xeda5, 0xb58, 0x156c, 0xb26d, 0x125c, 0x192c, 0x9a95, 0x1a94, 0x11b4a, 0x16aa, 0xad4, 0xd55b, 0x4ba, 0x125a, 0xb92b, 0x152a, 0x1694, 0x374a, 0x16aa, 0xeab5, 0x974, 0x14b6, 0xaa57, 0xa56, 0x1526, 0x8e95, 0xd54, 0x115aa, 0x156a, 0x96c, 0xd4ae, 0x14ae, 0xa4c, 0xbd26, 0x1b24, 0x1b54, 0x4d6a, 0x12da, 0xe95d, 0x95a, 0x149a, 0xba4b, 0x1a4a, 0x1aa4, 0x9b54, 0x16b4, 0x14ada, 0x12b6, 0x936, 0xd497, 0x1496, 0x164a, 0x96a5, 0xda4, 0x15b4, 0x4ab6, 0x126e, 0xc92f, 0x92e, 0xc96, 0xad4a, 0x1d4a, 0xd64, 0x75ac, 0x156c, 0x1126d, 0x125c, 0x192c, 0xda95, 0x1a94, 0x1b4a, 0xab55, 0xad4, 0x155a, 0x4a5d, 0xa5a, 0xf52b, 0x152a, 0x1694, 0xb6aa, 0x15aa, 0xab4, 0x74ba, 0x14b6, 0x10a57, 0xa4e, 0xd26, 0xce93, 0xd54, 0x15aa, 0x89b5, 0x96c, 0x14ae, 0x6a4e, 0x1a4c, 0xfd26, 0x1aa4, 0x1b52, 0xad6a, 0xada, 0x95c, 0x949d, 0x149a, 0x11a4b, 0x1a4a, 0x1aa4, 0xfb52, 0x16b4, 0xada, 0xa95b, 0x936, 0x1496, 0x7a4b, 0x154a, 0x116a5, 0xda4, 0x15ac, 0xaab6, 0xa6e, 0x92e, 0x8c97, 0xc96, 0x10d4a, 0x1d4a, 0xd54, 0xd56a, 0x155a, 0xa5c, 0xb92e, 0x192c, 0x1a94, 0x7d4a, 0x1b2a, 0xeb55, 0xad4, 0x14da, 0xaa5d, 0xa5a, 0x152a, 0x9a95, 0x1694, 0x156aa, 0x15aa, 0xab4, 0xd4ba, 0x14b6, 0xa56, 0xb517, 0xd26, 0xe52, 0x76a9, 0xdaa, 0xe5b5, 0x96c, 0x14ae, 0xaa2e, 0x1a2c, 0x1d16, 0x8d52, 0x1b52, 0x14d6a, 0xada, 0x55c, 0xd49d, 0x145a, 0x1a2a, 0xba95, 0x16a4, 0x1b52, 0x4b5a, 0xab6, 0xe55b, 0x8b6, 0x1456, 0xb62b, 0x152a, 0x16a4, 0x96d2, 0x15aa, 0xab6, 0x2537, 0x8ae, 0xcc57, 0xa56, 0xd2a, 0xad95, 0xb54, 0x156a, 0x6aad, 0x95c, 0xf4ae, 0x14ac, 0x1a94, 0xdd2a, 0x1b2a, 0xb54, 0x956a, 0x14da, 0x95c, 0x34ab, 0x151a, 0xfa8b, 0x1652, 0x16aa, 0xaad5, 0x6b4, 0x12ba, 0x695b, 0x936, 0xf517, 0xd16, 0xd52, 0xd6a9, 0xd6a, 0x5b4, 0x92b6, 0x12ae, 0x10a2e, 0x1a2c, 0x1c94, 0xfd4a, 0x1b52, 0xb6a, 0xa56d, 0x55c, 0x125c, 0x9a2d, 0x192a, 0xfa95, 0x1694, 0x16d2, 0xcb59, 0xab6, 0x53a, 0x925b, 0x1456, 0x1352b, 0x152a, 0x1694, 0xd6aa, 0x15aa, 0xab4, 0x94b7, 0x4ae, 0xa56, 0x752b, 0xd2a, 0xed95, 0xb54, 0x156a, 0xaa6d, 0x95c, 0x14ae, 0x8a55, 0x1a4c, 0x15d2a, 0x1aaa, 0xb54, 0xf56a, 0x12da, 0x95c, 0xb49b, 0x149a, 0x1a4a, 0x7b25, 0x16a8, 0x11ad5, 0x5b4, 0x12b6, 0xc957, 0x936, 0x1496, 0x964b, 0xd4a, 0x176a9, 0xd6a, 0x56c, 0xd2b6, 0x126e, 0x92e, 0xac96, 0x1c94, 0x1d4a, 0x6da9, 0xb58, 0x1156d, 0x55c, 0x125c, 0xd92d, 0x192a, 0x1a94, 0xbb4a, 0x16ca, 0xad4, 0x355b, 0x4ba, 0xf25b, 0x1256, 0x152a, 0xba95, 0xe94, 0x16aa, 0x6ad5, 0xab4, 0x114b6, 0x14ae, 0xa56, 0xd52a, 0x1d2a, 0xd94, 0x95aa, 0x156a, 0x96c, 0x54ae, 0x14ae, 0xea4d, 0x1a4c, 0x1d24, 0xbd55, 0xb54, 0x135a, 0x696d, 0x95a, 0xf49b, 0x149a, 0x1a4a, 0xdb25, 0x16a8, 0x1ad4, 0x92da, 0x12b6, 0x956, 0x7497, 0x1496, 0xf64b, 0xd4a, 0xda8, 0xb6b4, 0x156c, 0x12b6, 0x6937, 0x92e, 0xec96, 0x1a94, 0x1d4a, 0xcda5, 0xb58, 0x156c, 0x926d, 0x125c, 0x192c, 0x7c95, 0x1a94, 0xfb4a, 0x16aa, 0xad4, 0xb55a, 0x14ba, 0x125a, 0x792b, 0x152a, 0x11695, 0xe94, 0x16aa, 0xcad5, 0x9b4, 0x14b6, 0x8a57, 0xa56, 0x1526, 0x6e93, 0xd54, 0xf5aa, 0x156a, 0x96c, 0xb4ae, 0x14ac, 0x1a4c, 0x9d26, 0x1b24, 0x13b54, 0x1ad4, 0x135a, 0xc95d, 0x95a, 0x149a, 0xba4b, 0x1a4a, 0x17aa5, 0x16a8, 0x16d4, 0xeada, 0x12b6, 0x936, 0xb497, 0x1496, 0x164a, 0x96a5, 0xda4, 0x135b4, 0x156c, 0x12ae, 0xc92f, 0x92e, 0xc96, 0xad4a, 0x1d4a, 0xda4, 0x55aa, 0x156a, 0xea6d, 0x125c, 0x192c, 0xda95, 0x1a94, 0x1b4a, 0x8b55, 0xad4, 0x1155a, 0x14ba, 0xa5a, 0xd52b, 0x152a, 0x1694, 0xb6ca, 0x16aa, 0xab4, 0x74ba, 0x14b6, 0xea57, 0xa4e, 0xd26, 0xce93, 0xd54, 0x15aa, 0x69b5, 0x96c, 0x114ae, 0x149c, 0x1a4c, 0xdd26, 0x1aa4, 0x1b52, 0xad6a, 0xada, 0x95c, 0x74ad, 0x149a, 0xfa4b, 0x1a4a, 0x1aa4, 0xdb52, 0x16d4, 0xada, 0x695b, 0x936, 0x11497, 0x1496, 0x154a, 0xd6a5, 0xda4, 0x15ac, 0x8ab6, 0xa6e, 0x92e, 0x6c97, 0xc96, 0xed4a, 0x1d4a, 0xd64, 0xd5aa, 0x156a, 0xa6c, 0x992e, 0x192c, 0x13a95, 0x1a94, 0x1b4a, 0xeb55, 0xad4, 0x14da, 0xaa5d, 0xa5a, 0x152a, 0x7a95, 0x1692, 0x116aa, 0x15aa, 0xab4, 0xb4ba, 0x14b6, 0xa56, 0x9517, 0xd26, 0x12e53, 0xd52, 0xdaa, 0xc5b5, 0x96c, 0x14ae, 0xaa4e, 0x1a2c, 0x1d24, 0x7d52, 0x1b52, 0x10d6a, 0xada, 0x55c, 0xd49d, 0x149a, 0x1a2a, 0xbb15, 0x1aa4, 0x13b52, 0x16b2, 0xad6, 0xe55b, 0x936, 0x1456, 0xb62b, 0x154a, 0x16a4, 0x76d2, 0x15aa, 0xeab6, 0xa6e, 0x92e, 0xcc97, 0xa96, 0xd4a, 0x8e95, 0xd54, 0x1556a, 0x155a, 0xa6c, 0xd4ae, 0x152c, 0x1a94, 0xbd4a, 0x1b2a, 0xb54, 0x756a, 0x14da, 0xe95d, 0x956, 0x151a, 0xda8b, 0x1652, 0x16a8, 0x9ad5, 0xab4, 0x152ba, 0x12b6, 0xa56, 0xd517, 0xd16, 0xe52, 0xb6a9, 0xdaa, 0x5b4, 0x72b6, 0x12ae, 0xea2e, 0x1a2c, 0x1d14, 0xdd4a, 0x1b52, 0xb6a, 0x856d, 0x55c, 0x1725d, 0x145a, 0x1a2a, 0xfa95, 0x1694, 0x1b52, 0xab59, 0xab6, 0x55a, 0x725b, 0x1456, 0x1152b, 0x152a, 0x1694, 0xd6ca, 0x15aa, 0xab4, 0x94b6, 0x14ae, 0x12a57, 0xa56, 0xd2a, 0xed95, 0xb54, 0x156a, 0xaaad, 0x95c, 0x14ae, 0x8a55, 0x1a4c, 0x11d2a, 0x1b2a, 0xb54, 0xd56a, 0x12da, 0x95a, 0x94ab, 0x149a, 0x11a4b, 0x164a, 0x16a8, 0xfad5, 0x5b4, 0x12b6, 0xa95b, 0x936, 0x1496, 0x964b, 0xd4a, 0x116a9, 0xdaa, 0x5ac, 0xd2b6, 0x126e, 0x92e, 0x8c96, 0x1c94, 0x15d4a, 0x1b4a, 0xb68, 0xf56d, 0x55c, 0x125c, 0xb92d, 0x192a, 0x1a94, 0x7b4a, 0x16ca, 0x10ad5, 0xab6, 0x53a, 0xd25b, 0x1256, 0x152a, 0x9a95, 0x1694, 0x156aa, 0x15aa, 0xab4, 0xd4b6, 0x14ae, 0xa56, 0xb52a, 0x1d2a, 0xd94, 0x75aa, 0x156a, 0x1096d, 0x95c, 0x14ae, 0xca4d, 0x1a4c, 0x1d24, 0xbd54, 0x1b54, 0x136a, 0x496d, 0x95a, 0xf49b, 0x149a, 0x1a4a, 0xbb25, 0x16a8, 0x1ad4, 0x72da, 0x12b6, 0x10957, 0x92e, 0x1496, 0xd64b, 0xd4a, 0xea4, 0x96b4, 0x156c, 0x12b6, 0x4937, 0x92e, 0xec96, 0x1c94, 0x1d4a, 0xcda5, 0xb54, 0x156c, 0x92ad, 0x125c, 0x1192d, 0x192a, 0x1a94, 0xfb4a, 0x16ca, 0xad4, 0xb55a, 0x14ba, 0x125a, 0x752b, 0x152a, 0xf695, 0xe94, 0x16aa, 0xaad5, 0x9b4, 0x14b6, 0x8a57, 0xa56, 0x11526, 0x1d26, 0xd54, 0xd5aa, 0x156a, 0x96c, 0x94ae, 0x14ac, 0x17a4d, 0x1a4c, 0x1b24, 0xfd54, 0x1b54, 0xb5a, 0xa96d, 0x95a, 0x149a, 0x9a4b, 0x1a4a, 0x11b25, 0x16a4, 0x16d4, 0xcada, 0x12b6, 0x936, 0x9497, 0x1496, 0x1564b, 0xd4a, 0xda4, 0xf6b4, 0x156c, 0x12b6, 0xa937, 0x92e, 0xc96, 0x8d4a, 0x1d4a, 0x12da5, 0xb54, 0x156a, 0xca6d, 0x125c, 0x192c, 0xba95, 0x1a94, 0x1b4a, 0x4b55, 0xad4, 0xf55a, 0x14ba, 0xa5a, 0xd52b, 0x152a, 0x1694, 0x974a, 0x16aa, 0x12ad5}
	solar_1_1 = [...]int{1887, 0xec04c, 0xec23f, 0xec435, 0xec649, 0xec83e, 0xeca51, 0xecc46, 0xece3a, 0xed04d, 0xed242, 0xed436, 0xed64a, 0xed83f, 0xeda53, 0xedc48, 0xede3d, 0xee050, 0xee244, 0xee439, 0xee64d, 0xee842, 0xeea36, 0xeec4a, 0xeee3e, 0xef052, 0xef246, 0xef43a, 0xef64e, 0xef843, 0xefa37, 0xefc4b, 0xefe41, 0xf0054, 0xf0248, 0xf043c, 0xf0650, 0xf0845, 0xf0a38, 0xf0c4d, 0xf0e42, 0xf1037, 0xf124a, 0xf143e, 0xf1651, 0xf1846, 0xf1a3a, 0xf1c4e, 0xf1e44, 0xf2038, 0xf224b, 0xf243f, 0xf2653, 0xf2848, 0xf2a3b, 0xf2c4f, 0xf2e45, 0xf3039, 0xf324d, 0xf3442, 0xf3636, 0xf384a, 0xf3a3d, 0xf3c51, 0xf3e46, 0xf403b, 0xf424e, 0xf4443, 0xf4638, 0xf484c, 0xf4a3f, 0xf4c52, 0xf4e48, 0xf503c, 0xf524f, 0xf5445, 0xf5639, 0xf584d, 0xf5a42, 0xf5c35, 0xf5e49, 0xf603e, 0xf6251, 0xf6446, 0xf663b, 0xf684f, 0xf6a43, 0xf6c37, 0xf6e4b, 0xf703f, 0xf7252, 0xf7447, 0xf763c, 0xf7850, 0xf7a45, 0xf7c39, 0xf7e4d, 0xf8042, 0xf8254, 0xf8449, 0xf863d, 0xf8851, 0xf8a46, 0xf8c3b, 0xf8e4f, 0xf9044, 0xf9237, 0xf944a, 0xf963f, 0xf9853, 0xf9a47, 0xf9c3c, 0xf9e50, 0xfa045, 0xfa238, 0xfa44c, 0xfa641, 0xfa836, 0xfaa49, 0xfac3d, 0xfae52, 0xfb047, 0xfb23a, 0xfb44e, 0xfb643, 0xfb837, 0xfba4a, 0xfbc3f, 0xfbe53, 0xfc048, 0xfc23c, 0xfc450, 0xfc645, 0xfc839, 0xfca4c, 0xfcc41, 0xfce36, 0xfd04a, 0xfd23d, 0xfd451, 0xfd646, 0xfd83a, 0xfda4d, 0xfdc43, 0xfde37, 0xfe04b, 0xfe23f, 0xfe453, 0xfe648, 0xfe83c, 0xfea4f, 0xfec44, 0xfee38, 0xff04c, 0xff241, 0xff436, 0xff64a, 0xff83e, 0xffa51, 0xffc46, 0xffe3a, 0x10004e, 0x100242, 0x100437, 0x10064b, 0x100841, 0x100a53, 0x100c48, 0x100e3c, 0x10104f, 0x101244, 0x101438, 0x10164c, 0x101842, 0x101a35, 0x101c49, 0x101e3d, 0x102051, 0x102245, 0x10243a, 0x10264e, 0x102843, 0x102a37, 0x102c4b, 0x102e3f, 0x103053, 0x103247, 0x10343b, 0x10364f, 0x103845, 0x103a38, 0x103c4c, 0x103e42, 0x104036, 0x104249, 0x10443d, 0x104651, 0x104846, 0x104a3a, 0x104c4e, 0x104e43, 0x105038, 0x10524a, 0x10543e, 0x105652, 0x105847, 0x105a3b, 0x105c4f, 0x105e45, 0x106039, 0x10624c, 0x106441, 0x106635, 0x106849, 0x106a3d, 0x106c51, 0x106e47, 0x10703c, 0x10724f, 0x107444, 0x107638, 0x10784c, 0x107a3f, 0x107c53, 0x107e48, 0x10803d, 0x108250, 0x108446, 0x10863a, 0x10884e, 0x108a42, 0x108c36, 0x108e4a, 0x10903e, 0x109251, 0x109447, 0x10963b, 0x10984f, 0x109a43, 0x109c37, 0x109e4b, 0x10a041, 0x10a253, 0x10a448, 0x10a63d, 0x10a851, 0x10aa45, 0x10ac39, 0x10ae4d, 0x10b042, 0x10b236, 0x10b44a, 0x10b63e, 0x10b852, 0x10ba47, 0x10bc3b, 0x10be4f, 0x10c044, 0x10c237, 0x10c44b, 0x10c641, 0x10c854, 0x10ca48, 0x10cc3d, 0x10ce50, 0x10d045, 0x10d239, 0x10d44c, 0x10d642, 0x10d837, 0x10da4a, 0x10dc3e, 0x10de52, 0x10e047, 0x10e23a, 0x10e44e, 0x10e643, 0x10e838, 0x10ea4b, 0x10ec41, 0x10ee54, 0x10f049, 0x10f23c, 0x10f450, 0x10f645, 0x10f839, 0x10fa4c, 0x10fc42, 0x10fe37, 0x11004b, 0x11023e, 0x110452, 0x110647, 0x11083b, 0x110a4e, 0x110c43, 0x110e38, 0x11104c, 0x11123f, 0x111435, 0x111648, 0x11183c, 0x111a4f, 0x111c45, 0x111e39, 0x11204d, 0x112242, 0x112436, 0x11264a, 0x11283e, 0x112a51, 0x112c46, 0x112e3b, 0x11304f, 0x113244, 0x113439, 0x11364d, 0x113842, 0x113a54, 0x113c49, 0x113e3d, 0x114051, 0x114246, 0x11443a, 0x11464e, 0x114844, 0x114a37, 0x114c4a, 0x114e3e, 0x115052, 0x115247, 0x11543c, 0x115650, 0x115845, 0x115a38, 0x115c4c, 0x115e41, 0x116054, 0x116248, 0x11643d, 0x116651, 0x116847, 0x116a3a, 0x116c4e, 0x116e43, 0x117037, 0x11724a, 0x11743e, 0x117652, 0x117848, 0x117a3c, 0x117c50, 0x117e45, 0x118039, 0x11824c, 0x118441, 0x118654, 0x118849, 0x118a3d, 0x118c51, 0x118e46, 0x11903b, 0x11924d, 0x119442, 0x119637, 0x11984b, 0x119a3e, 0x119c52, 0x119e48, 0x11a03c, 0x11a24f, 0x11a444, 0x11a638, 0x11a84c, 0x11aa3f, 0x11ac35, 0x11ae49, 0x11b03e, 0x11b251, 0x11b446, 0x11b63a, 0x11b84e, 0x11ba42, 0x11bc36, 0x11be4b, 0x11c03f, 0x11c252, 0x11c448, 0x11c63c, 0x11c84f, 0x11ca43, 0x11cc38, 0x11ce4c, 0x11d042, 0x11d235, 0x11d449, 0x11d63d, 0x11d851, 0x11da45, 0x11dc39, 0x11de4d, 0x11e043, 0x11e236, 0x11e44b, 0x11e63f, 0x11e853, 0x11ea47, 0x11ec3b, 0x11ee4f, 0x11f044, 0x11f238, 0x11f44c, 0x11f641, 0x11f836, 0x11fa4a, 0x11fc3e, 0x11fe51, 0x120047, 0x12023a, 0x12044e, 0x120644, 0x120838, 0x120a4b, 0x120c41, 0x120e53, 0x121048, 0x12123c, 0x121450, 0x121645, 0x12183a, 0x121a4d, 0x121c42, 0x121e55, 0x12204a, 0x12223d, 0x122451, 0x122646, 0x12283b, 0x122a4e, 0x122c44, 0x122e38, 0x12304c, 0x12323f, 0x123453, 0x123648, 0x12383c, 0x123a4f, 0x123c45, 0x123e3a, 0x12404e, 0x124242, 0x124436, 0x124649, 0x12483e, 0x124a51, 0x124c46, 0x124e3b, 0x12504f, 0x125243, 0x125437, 0x12564b, 0x12583f, 0x125a52, 0x125c48, 0x125e3c, 0x126051, 0x126245, 0x126439, 0x12664d, 0x126842, 0x126a35, 0x126c49, 0x126e3e, 0x127052, 0x127246, 0x12743b, 0x12764f, 0x127844, 0x127a37, 0x127c4b, 0x127e3f, 0x128053, 0x128248, 0x12843c, 0x128650, 0x128846, 0x128a38, 0x128c4c, 0x128e41, 0x129036, 0x129249, 0x12943e, 0x129652, 0x129847, 0x129a3a, 0x129c4e, 0x129e43, 0x12a037, 0x12a24a, 0x12a43f, 0x12a653, 0x12a849, 0x12aa3c, 0x12ac50, 0x12ae45, 0x12b039, 0x12b24c, 0x12b441, 0x12b636, 0x12b84a, 0x12ba3e, 0x12bc52, 0x12be47, 0x12c03b, 0x12c24d, 0x12c443, 0x12c637, 0x12c84b, 0x12ca3f, 0x12cc53, 0x12ce48, 0x12d03c, 0x12d24f, 0x12d444, 0x12d639, 0x12d84d, 0x12da41, 0x12dc36, 0x12de4a, 0x12e03e, 0x12e251, 0x12e446, 0x12e63a, 0x12e84e, 0x12ea43, 0x12ec37, 0x12ee4b, 0x12f041, 0x12f253, 0x12f448, 0x12f63c, 0x12f850, 0x12fa44, 0x12fc39, 0x12fe4d, 0x130042, 0x130236, 0x130449, 0x13063d, 0x130851, 0x130a45, 0x130c3a, 0x130e4e, 0x131044, 0x131237, 0x13144b, 0x13163f, 0x131853, 0x131a47, 0x131c3b, 0x131e4f, 0x132045, 0x132239, 0x13244d, 0x132642, 0x132836, 0x132a49, 0x132c3d, 0x132e51, 0x133046, 0x13323a, 0x13344e, 0x133644, 0x133838, 0x133a4b, 0x133c3f, 0x133e52, 0x134048, 0x13423b, 0x13444f, 0x134645, 0x134839, 0x134a4c, 0x134c41, 0x134e35, 0x135049, 0x13523d, 0x135451, 0x135646, 0x13583b, 0x135a4e, 0x135c43, 0x135e37, 0x13604b, 0x13623e, 0x136452, 0x136648, 0x13683c, 0x136a4f, 0x136c45, 0x136e39, 0x13704d, 0x137241, 0x137435, 0x137649, 0x13783e, 0x137a51, 0x137c46, 0x137e3b, 0x13804e, 0x138242, 0x138437, 0x13864a, 0x13883f, 0x138a53, 0x138c49, 0x138e3d, 0x139051, 0x139245, 0x139439, 0x13964d, 0x139842, 0x139a36, 0x139c4a, 0x139e3f, 0x13a053, 0x13a247, 0x13a43b, 0x13a64f, 0x13a844, 0x13aa37, 0x13ac4b, 0x13ae41, 0x13b036, 0x13b249, 0x13b43d, 0x13b651, 0x13b846, 0x13ba39, 0x13bc4d, 0x13be42, 0x13c037, 0x13c24a, 0x13c43e, 0x13c652, 0x13c847, 0x13ca3a, 0x13cc4e, 0x13ce44, 0x13d038, 0x13d24b, 0x13d441, 0x13d635, 0x13d849, 0x13da3c, 0x13dc50, 0x13de45, 0x13e03a, 0x13e24d, 0x13e442, 0x13e637, 0x13e84b, 0x13ea3e, 0x13ec52, 0x13ee47, 0x13f03b, 0x13f24e, 0x13f444, 0x13f638, 0x13f84c, 0x13fa41, 0x13fc53, 0x13fe48, 0x14003c, 0x14024f, 0x140445, 0x14063a, 0x14084e, 0x140a42, 0x140c36, 0x140e4a, 0x14103e, 0x141251, 0x141446, 0x14163b, 0x14184f, 0x141a44, 0x141c38, 0x141e4c, 0x142041, 0x142253, 0x142448, 0x14263c, 0x142850, 0x142a45, 0x142c39, 0x142e4e, 0x143043, 0x143236, 0x14344a, 0x14363e, 0x143852, 0x143a46, 0x143c3b, 0x143e4f, 0x144044, 0x144238, 0x14444b, 0x14463f, 0x144853, 0x144a48, 0x144c3c, 0x144e50, 0x145046, 0x14523a, 0x14544e, 0x145643, 0x145837, 0x145a4a, 0x145c3e, 0x145e53, 0x146048, 0x14623c, 0x146450, 0x146645, 0x146839, 0x146a4c, 0x146c41, 0x146e54, 0x14704a, 0x14723d, 0x147451, 0x147647, 0x14783b, 0x147a4d, 0x147c42, 0x147e37, 0x14804b, 0x14823e, 0x148453, 0x148648, 0x14883c, 0x148a4f, 0x148c44, 0x148e38, 0x14904c, 0x149241, 0x149435, 0x14964a, 0x14983e, 0x149a51, 0x149c46, 0x149e3a, 0x14a04e, 0x14a242, 0x14a437, 0x14a64b, 0x14a83f, 0x14aa52, 0x14ac48, 0x14ae3c, 0x14b04f, 0x14b244, 0x14b438, 0x14b64c, 0x14b842, 0x14ba35, 0x14bc49, 0x14be3e, 0x14c051, 0x14c245, 0x14c439, 0x14c64e, 0x14c843, 0x14ca37, 0x14cc4b, 0x14ce3f, 0x14d053, 0x14d247, 0x14d43b, 0x14d64f, 0x14d844, 0x14da38, 0x14dc4c, 0x14de42, 0x14e036, 0x14e249, 0x14e43d, 0x14e651, 0x14e846, 0x14ea39, 0x14ec4d, 0x14ee43, 0x14f038, 0x14f24b, 0x14f43f, 0x14f652, 0x14f847, 0x14fa3b, 0x14fc4f, 0x14fe44, 0x150039, 0x15024c, 0x150441, 0x150635, 0x150849, 0x150a3c, 0x150c50, 0x150e46, 0x15103a, 0x15124d, 0x151443, 0x151637, 0x15184b, 0x151a3f, 0x151c53, 0x151e48, 0x15203d, 0x152250, 0x152445, 0x15263a, 0x15284e, 0x152a42, 0x152c36, 0x152e4a, 0x15303e, 0x153251, 0x153447, 0x15363b, 0x15384f, 0x153a43, 0x153c37, 0x153e4b, 0x15403f, 0x154252, 0x154448, 0x15463d, 0x154851, 0x154a45, 0x154c39, 0x154e4d, 0x155042, 0x155254, 0x155449, 0x15563e, 0x155852, 0x155a47, 0x155c3b, 0x155e4f, 0x156044, 0x156237, 0x15644b, 0x15663f, 0x156853, 0x156a48, 0x156c3d, 0x156e51, 0x157046, 0x157239, 0x15744c, 0x157642, 0x157836, 0x157a49, 0x157c3e, 0x157e52, 0x158047, 0x15823a, 0x15844e, 0x158643, 0x158838, 0x158a4b, 0x158c3f, 0x158e53, 0x159049, 0x15923c, 0x159450, 0x159645, 0x159839, 0x159a4c, 0x159c42, 0x159e36, 0x15a04a, 0x15a23e, 0x15a452, 0x15a647, 0x15a83b, 0x15aa4e, 0x15ac43, 0x15ae38, 0x15b04c, 0x15b23f, 0x15b453, 0x15b648, 0x15b83c, 0x15ba4f, 0x15bc44, 0x15be39, 0x15c04d, 0x15c242, 0x15c436, 0x15c64a, 0x15c83e, 0x15ca51, 0x15cc46, 0x15ce3a, 0x15d04e, 0x15d243, 0x15d438, 0x15d64c, 0x15d841, 0x15da53, 0x15dc48, 0x15de3c, 0x15e050, 0x15e244, 0x15e439, 0x15e64d, 0x15e843, 0x15ea36, 0x15ec4a, 0x15ee3e, 0x15f051, 0x15f246, 0x15f43a, 0x15f64e, 0x15f844, 0x15fa37, 0x15fc4b, 0x15fe3f, 0x160053, 0x160247, 0x16043c, 0x160650, 0x160845, 0x160a39, 0x160c4d, 0x160e42, 0x161036, 0x161249, 0x16143d, 0x161651, 0x161847, 0x161a3a, 0x161c4e, 0x161e44, 0x162038, 0x16224b, 0x16243f, 0x162653, 0x162848, 0x162a3c, 0x162c50, 0x162e45, 0x16303a, 0x16324c, 0x163441, 0x163635, 0x163849, 0x163a3d, 0x163c51, 0x163e47, 0x16403b, 0x16424e, 0x164443, 0x164637, 0x16484b, 0x164a3e, 0x164c52, 0x164e48, 0x16503d, 0x165250, 0x165445, 0x165639, 0x16584d, 0x165a41, 0x165c35, 0x165e49, 0x16603e, 0x166251, 0x166447, 0x16663b, 0x16684f, 0x166a43, 0x166c37, 0x166e4b, 0x16703f, 0x167252, 0x167448, 0x16763c, 0x167850, 0x167a44, 0x167c38, 0x167e4c, 0x168042, 0x168235, 0x168449, 0x16863e, 0x168852, 0x168a46, 0x168c3a, 0x168e4e, 0x169043, 0x169236, 0x16944b, 0x16963f, 0x169853, 0x169a48, 0x169c3c, 0x169e50, 0x16a045, 0x16a238, 0x16a44c, 0x16a642, 0x16a836, 0x16aa4a, 0x16ac3f, 0x16ae52, 0x16b047, 0x16b23a, 0x16b44e, 0x16b644, 0x16b839, 0x16ba4c, 0x16bc41, 0x16be54, 0x16c049, 0x16c23c, 0x16c450, 0x16c645, 0x16c83a, 0x16ca4d, 0x16cc43, 0x16ce37, 0x16d04b, 0x16d23e, 0x16d452, 0x16d647, 0x16d83b, 0x16da4e, 0x16dc44, 0x16de38, 0x16e04c, 0x16e241, 0x16e454, 0x16e649, 0x16e83d, 0x16ea50, 0x16ec45, 0x16ee3a, 0x16f04e, 0x16f242, 0x16f437, 0x16f64a, 0x16f83e, 0x16fa51, 0x16fc47, 0x16fe3b, 0x17004f, 0x170244, 0x170438, 0x17064c, 0x170841, 0x170a53, 0x170c48, 0x170e3c, 0x171051, 0x171245, 0x17143a, 0x17164e, 0x171843, 0x171a36, 0x171c4a, 0x171e3e, 0x172052, 0x172247, 0x17243b, 0x17264f, 0x172845, 0x172a38, 0x172c4b, 0x172e3f, 0x173053, 0x173248, 0x17343c, 0x173651, 0x173846, 0x173a39, 0x173c4d, 0x173e42, 0x174036, 0x174249, 0x17443e, 0x174652, 0x174848, 0x174a3b, 0x174c4f, 0x174e44, 0x175038, 0x17524b, 0x17543f, 0x175653, 0x175849, 0x175a3c, 0x175c50, 0x175e46, 0x17603a, 0x17624c, 0x176442, 0x176636, 0x17684a, 0x176a3e, 0x176c52, 0x176e47, 0x17703c, 0x17724f, 0x177444, 0x177638, 0x17784c, 0x177a41, 0x177c36, 0x177e4a, 0x17803e, 0x178251, 0x178446, 0x17863a, 0x17884e, 0x178a42, 0x178c37, 0x178e4b, 0x179041, 0x179253, 0x179448, 0x17963c, 0x179850, 0x179a44, 0x179c38, 0x179e4c, 0x17a042, 0x17a236, 0x17a44a, 0x17a63e, 0x17a851, 0x17aa45, 0x17ac3a, 0x17ae4e, 0x17b043, 0x17b237, 0x17b44b, 0x17b63f, 0x17b853, 0x17ba47, 0x17bc3b, 0x17be4f, 0x17c045, 0x17c238, 0x17c44c, 0x17c642, 0x17c836, 0x17ca49, 0x17cc3d, 0x17ce51, 0x17d046, 0x17d23a, 0x17d44e, 0x17d643, 0x17d838, 0x17da4b, 0x17dc3f, 0x17de52, 0x17e048, 0x17e23b, 0x17e44f, 0x17e645, 0x17e839, 0x17ea4c, 0x17ec41, 0x17ee54, 0x17f049, 0x17f23c, 0x17f450, 0x17f646, 0x17f83b, 0x17fa4e, 0x17fc43, 0x17fe37, 0x18004b, 0x18023e, 0x180452, 0x180647, 0x18083c, 0x180a4f, 0x180c45, 0x180e39, 0x18104d, 0x181241, 0x181454, 0x181649, 0x18183d, 0x181a50, 0x181c46, 0x181e3b, 0x18204f, 0x182243, 0x182437, 0x18264a, 0x18283f, 0x182a52, 0x182c47, 0x182e3c, 0x183050, 0x183244, 0x183438, 0x18364c, 0x183841, 0x183a36, 0x183c4a, 0x183e3e, 0x184052, 0x184247, 0x18443b, 0x18464f, 0x184844, 0x184a37, 0x184c4b, 0x184e41}
)
