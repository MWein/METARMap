package metar

// Airports must be in the same order that you wish to encode them to the Shift Registers

var Airports = []string{
	"KSET",
	"KSTL",
	"KMQB",

	"K00C",
	"K00F",
	"K00M",
	"K00R",
	"K00V",
	"K01",
	"K01G",
	"K01M",
	"K02A",
	"K02C",
	"K02G",
	"K03",
	"K03B",
	"K03D",
	"K04A",
	"K04G",
	"K04M",
	"K04Y",
	"K05C",
	"K05D",
	"K05U",
	"K06A",
	"K06C",
	"K06D",
	"K06M",
	"K06U",
	"K07",
	"K07A",
	"K07F",
	"K07R",
	"K07V",
	"K08",
	"K08A",
	"K08C",
	"K08D",
	"K08K",
	"K08M",
	"K09A",
	"K09J",
	"K09K",
	"K09M",
	"K09R",
	"K0A2",
	"K0A3",
	"K0A4",
	"K0A7",
	"K0A8",
	"K0A9",
	"K0B1",
	"K0B4",
	"K0B5",
	"K0B8",
	"K0C0",
	"K0C4",
	"K0D8",
	"K0E0",
	"K0E8",
	"K0F2",
	"K0F4",
	"K0F7",
	"K0F9",
	"K0G3",
	"K0G6",
	"K0G7",
	"K0H1",
	"K0I8",
	"K0J4",
	"K0J6",
	"K0J9",
	"K0K7",
	"K0L7",
	"K0L9",
	"K0M0",
	"K0M1",
	"K0M2",
	"K0M3",
	"K0M4",
	"K0M5",
	"K0M8",
	"K0M9",
	"K0O2",
	"K0Q5",
	"K0R0",
	"K0R1",
	"K0R3",
	"K0R4",
	"K0R5",
	"K0R6",
	"K0S0",
	"K0S7",
	"K0S9",
	"K0V3",
	"K0V4",
	"K0V7",
	"K0VG",
	"K0W3",
	"K10C",
	"K10G",
	"K10U",
	"K11",
	"K11A",
	"K11R",
	"K11V",
	"K12D",
	"K12G",
	"K12J",
	"K12K",
	"K12N",
	"K12V",
	"K12Y",
	"K13",
	"K13C",
	"K13K",
	"K14A",
	"K14F",
	"K14G",
	"K14J",
	"K14M",
	"K14Y",
	"K15F",
	"K15J",
	"K15M",
	"K16D",
	"K16G",
	"K16J",
	"K17",
	"K17G",
	"K17J",
	"K17K",
	"K17M",
	"K17N",
	"K17V",
	"K18A",
	"K18I",
	"K18V",
	"K19A",
	"K19M",
	"K19N",
	"K19S",
	"K1A0",
	"K1A3",
	"K1A4",
	"K1A5",
	"K1A6",
	"K1A7",
	"K1A9",
	"K1B0",
	"K1B1",
	"K1B2",
	"K1B6",
	"K1B9",
	"K1BT",
	"K1C1",
	"K1C2",
	"K1C5",
	"K1D1",
	"K1D3",
	"K1D7",
	"K1D8",
	"K1F0",
	"K1F4",
	"K1F5",
	"K1G0",
	"K1G1",
	"K1G3",
	"K1G4",
	"K1G5",
	"K1H0",
	"K1H2",
	"K1H3",
	"K1H5",
	"K1H8",
	"K1I5",
	"K1J0",
	"K1J6",
	"K1K2",
	"K1K4",
	"K1K7",
	"K1K9",
	"K1L0",
	"K1L1",
	"K1L3",
	"K1L7",
	"K1L8",
	"K1L9",
	"K1M2",
	"K1M4",
	"K1M5",
	"K1M9",
	"K1MO",
	"K1N1",
	"K1N7",
	"K1O1",
	"K1O2",
	"K1O3",
	"K1O4",
	"K1O5",
	"K1O8",
	"K1Q2",
	"K1Q4",
}