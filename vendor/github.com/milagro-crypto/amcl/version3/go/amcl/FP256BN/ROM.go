/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

/* Fixed Data in ROM - Field and Curve parameters */

package FP256BN

// Base Bits= 56
var Modulus= [...]Chunk {0x292DDBAED33013,0x65FB12980A82D3,0x5EEE71A49F0CDC,0xFFFCF0CD46E5F2,0xFFFFFFFF}
var R2modp= [...]Chunk {0xEDE336303B9F8B,0x92FFEE9FEC54E8,0x13C1C063C55F79,0xA12F2EAC0123FA,0x8E559B2A}
const MConst Chunk=0x6C964E0537E5E5

const CURVE_A int= 0
const CURVE_B_I int= 3
var CURVE_B= [...]Chunk {0x3,0x0,0x0,0x0,0x0}
var CURVE_Order= [...]Chunk {0x2D536CD10B500D,0x65FB1299921AF6,0x5EEE71A49E0CDC,0xFFFCF0CD46E5F2,0xFFFFFFFF}
var CURVE_Gx= [...]Chunk {0x1,0x0,0x0,0x0,0x0}
var CURVE_Gy= [...]Chunk {0x2,0x0,0x0,0x0,0x0}

var Fra= [...]Chunk {0x760328AF943106,0x71511E3AB28F74,0x8DDB0867CF39A1,0xCA786F352D1A6E,0x3D617662}
var Frb= [...]Chunk {0xB32AB2FF3EFF0D,0xF4A9F45D57F35E,0xD113693CCFD33A,0x3584819819CB83,0xC29E899D}
var CURVE_Bnx= [...]Chunk {0x82F5C030B0A801,0x68,0x0,0x0,0x0}
var CURVE_Cof= [...]Chunk {0x1,0x0,0x0,0x0,0x0}
var CURVE_Cru= [...]Chunk {0x1C0A24A3A1B807,0xD79DF1932D1EDB,0x40921018659BCD,0x13988E1,0x0}
var CURVE_Pxa= [...]Chunk {0x2616B689C09EFB,0x539A12BF843CD2,0x577C28913ACE1C,0xB4C96C2028560F,0xFE0C3350}
var CURVE_Pxb= [...]Chunk {0x69ED34A37E6A2B,0x78E287D03589D2,0xC637D813B924DD,0x738AC054DB5AE1,0x4EA66057}
var CURVE_Pya= [...]Chunk {0x9B481BEDC27FF,0x24758D615848E9,0x75124E3E51EFCB,0xC542A3B376770D,0x702046E7}
var CURVE_Pyb= [...]Chunk {0x1281114AAD049B,0xBE80821A98B3E0,0x49297EB29F8B4C,0xD388C29042EEA6,0x554E3BC}
var CURVE_W=[2][5]Chunk {{0xF0036E1B054003,0xFFFFFFFE78663A,0xFFFF,0x0,0x0},{0x5EB8061615001,0xD1,0x0,0x0,0x0}}
var CURVE_SB=[2][2][5]Chunk {{{0xF5EEEE7C669004,0xFFFFFFFE78670B,0xFFFF,0x0,0x0},{0x5EB8061615001,0xD1,0x0,0x0,0x0}},{{0x5EB8061615001,0xD1,0x0,0x0,0x0},{0x3D4FFEB606100A,0x65FB129B19B4BB,0x5EEE71A49D0CDC,0xFFFCF0CD46E5F2,0xFFFFFFFF}}}
var CURVE_WB=[4][5]Chunk {{0x20678F0D30A800,0x55555554D2CC10,0x5555,0x0,0x0},{0xD6764C0D7DC805,0x8FBEA10BC3AD1A,0x806160104467DE,0xD105EB,0x0},{0xACB6061F173803,0x47DF5085E1D6C1,0xC030B0082233EF,0x6882F5,0x0},{0x26530F6E91F801,0x55555554D2CCE1,0x5555,0x0,0x0}}
var CURVE_BB=[4][4][5]Chunk {{{0xAA5DACA05AA80D,0x65FB1299921A8D,0x5EEE71A49E0CDC,0xFFFCF0CD46E5F2,0xFFFFFFFF},{0xAA5DACA05AA80C,0x65FB1299921A8D,0x5EEE71A49E0CDC,0xFFFCF0CD46E5F2,0xFFFFFFFF},{0xAA5DACA05AA80C,0x65FB1299921A8D,0x5EEE71A49E0CDC,0xFFFCF0CD46E5F2,0xFFFFFFFF},{0x5EB8061615002,0xD1,0x0,0x0,0x0}},{{0x5EB8061615001,0xD1,0x0,0x0,0x0},{0xAA5DACA05AA80C,0x65FB1299921A8D,0x5EEE71A49E0CDC,0xFFFCF0CD46E5F2,0xFFFFFFFF},{0xAA5DACA05AA80D,0x65FB1299921A8D,0x5EEE71A49E0CDC,0xFFFCF0CD46E5F2,0xFFFFFFFF},{0xAA5DACA05AA80C,0x65FB1299921A8D,0x5EEE71A49E0CDC,0xFFFCF0CD46E5F2,0xFFFFFFFF}},{{0x5EB8061615002,0xD1,0x0,0x0,0x0},{0x5EB8061615001,0xD1,0x0,0x0,0x0},{0x5EB8061615001,0xD1,0x0,0x0,0x0},{0x5EB8061615001,0xD1,0x0,0x0,0x0}},{{0x82F5C030B0A802,0x68,0x0,0x0,0x0},{0xBD700C2C2A002,0x1A2,0x0,0x0,0x0},{0x2767EC6FAA000A,0x65FB1299921A25,0x5EEE71A49E0CDC,0xFFFCF0CD46E5F2,0xFFFFFFFF},{0x82F5C030B0A802,0x68,0x0,0x0,0x0}}}

const USE_GLV bool=true
const USE_GS_G2 bool=true
const USE_GS_GT bool=true	
const GT_STRONG bool=false
