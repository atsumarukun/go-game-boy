package cpu

// Export guts for testing.

type Registers = registers

var ReadAF = (*registers).readAF
var ReadBC = (*registers).readBC
var ReadDE = (*registers).readDE
var ReadHL = (*registers).readHL

var WriteAF = (*registers).writeAF
var WriteBC = (*registers).writeBC
var WriteDE = (*registers).writeDE
var WriteHL = (*registers).writeHL

var GetZF = (*registers).getZF
var GetNF = (*registers).getNF
var GetHF = (*registers).getHF
var GetCF = (*registers).getCF

var SetZF = (*registers).setZF
var SetNF = (*registers).setNF
var SetHF = (*registers).setHF
var SetCF = (*registers).setCF
