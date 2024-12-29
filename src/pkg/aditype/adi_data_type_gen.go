// DO NOT EDIT
// Code generated by ./cmd/codegen/codegen

// This file contains CURRENTLY valid constants.
// It is not a full list of all possible historic and/or future values.

package aditype

var (
	CreditList                                DataType = "CreditList"                                // CreditList = a comma-delimited list where each list item is either: A member of the Credit enumeration. A member of the Credit enumeration followed by a colon and an ampersand-delimited list of members of the QSL_Medium enumeration. For example IOTA,WAS:LOTW&CARD,DXCC:CARD
	SponsoredAwardList                        DataType = "SponsoredAwardList"                        // SponsoredAwardList = a comma-delimited list of members of the Sponsored_Award enumeration
	Boolean                                   DataType = "Boolean"                                   // Boolean = if True, the single ASCII character Y or y if False, the single ASCII character N or n
	Digit                                     DataType = "Digit"                                     // Digit = an ASCII character whose code lies in the range of 48 through 57, inclusive
	Integer                                   DataType = "Integer"                                   // Integer = a sequence of one or more Digits representing a decimal integer, optionally preceded by a minus sign (ASCII code 45). Leading zeroes are allowed.
	Number                                    DataType = "Number"                                    // Number = a sequence of one or more Digits representing a decimal number, optionally preceded by a minus sign (ASCII code 45) and optionally including a single decimal point (ASCII code 46)
	PositiveInteger                           DataType = "PositiveInteger"                           // PositiveInteger = an unsigned sequence of one or more Digits representing a decimal integer that has a value greater than 0. Leading zeroes are allowed.
	Character                                 DataType = "Character"                                 // Character = an ASCII character whose code lies in the range of 32 through 126, inclusive
	IntlCharacter                             DataType = "IntlCharacter"                             // IntlCharacter = a Unicode character (encoded with UTF-8) excluding line break CR (code 13) and LF (code 10) characters
	Date                                      DataType = "Date"                                      // Date = 8 Digits representing a UTC date in YYYYMMDD format, where YYYY is a 4-Digit year specifier, where 1930 <= YYYY MM is a 2-Digit month specifier, where 1 <= MM <= 12 [use leading zeroes] DD is a 2-Digit day specifier, where 1 <= DD <= DaysInMonth(MM) [use leading zeroes]
	Time                                      DataType = "Time"                                      // Time = 6 Digits representing a UTC time in HHMMSS format or 4 Digits representing a time in HHMM format, where HH is a 2-Digit hour specifier, where 0 <= HH <= 23 [use leading zeroes] MM is a 2-Digit minute specifier, where 0 <= MM <= 59 [use leading zeroes] SS is a 2-Digit second specifier, where 0 <= SS <= 59 [use leading zeroes]
	IOTARefNo                                 DataType = "IOTARefNo"                                 // IOTARefNo = IOTA designator, in format CC-XXX, where CC is a member of the Continent enumeration XXX is the island group designator, where 1 <= XXX <= 999 [use leading zeroes]
	String                                    DataType = "String"                                    // String = a sequence of Characters
	IntlString                                DataType = "IntlString"                                // IntlString = a sequence of International Characters. Fields of type IntlString must only be used in ADX files
	MultilineString                           DataType = "MultilineString"                           // MultilineString = a sequence of Characters and line-breaks, where a line break is an ASCII CR (code 13) followed immediately by an ASCII LF (code 10)
	IntlMultilineString                       DataType = "IntlMultilineString"                       // IntlMultilineString = a sequence of International Characters and line breaks. Fields of type IntlMultilineString must only be used in ADX files
	Enumeration                               DataType = "Enumeration"                               // Enumeration = an explicit list of legal case-insensitive values represented in ASCII set forth in set notation, e.g. {A, B, C, D}, or defined in a table, from which a single value may be selected.
	GridSquare                                DataType = "GridSquare"                                // GridSquare = a case-insensitive 2-character, 4-character, 6-character, or 8-character Maidenhead locator. Specific fields impose additional restrictions on the number of characters; see the field descriptions for the allowed numbers of characters.
	GridSquareExt                             DataType = "GridSquareExt"                             // GridSquareExt = For a 10-character Maidenhead locator, contains characters 9 and 10. For a 12-character Maidenhead locator, contains characters 9, 10, 11 and 12. Characters 9 and 10 are case-insensitive ASCII letters in the range A-X. Characters 11 and 12 are Digits in the range 0-9.
	GridSquareList                            DataType = "GridSquareList"                            // GridSquareList = a comma-delimited list of GridSquare items
	Location                                  DataType = "Location"                                  // Location = a sequence of 11 characters representing a latitude or longitude in XDDD MM.MMM format, where X is a directional Character from the set {E, W, N, S} DDD is a 3-Digit degrees specifier, where 0 <= DDD <= 180 [use leading zeroes] There is a single space character in between DDD and MM.MMM MM.MMM is an unsigned Number minutes specifier with its decimal point in the third position, where 00.000 <= MM.MMM <= 59.999 [use leading and trailing zeroes]
	POTARef                                   DataType = "POTARef"                                   // POTARef = a sequence of case-insensitive Characters representing a Parks on the Air park reference in the form xxxx-nnnnn[@yyyyyy] comprising 6 to 17 characters where: xxxx is the POTA national program and is 1 to 4 characters in length, typically the default callsign prefix of the national program (rather than the DX entity) nnnnn represents the unique number within the national program and is either 4 or 5 characters in length (use the exact format listed on the POTA website) yyyyyy **Optional** is the 4 to 6 character ISO 3166-2 code to differentiate which state/province/prefecture/primary administration location the contact represents, in the case that the park reference spans more than one location (such as a trail). Examples of the POTARef Data Type: ReferenceLocation K-5033Golden Hill State Forest K-10000 5-digit park numbers are reserved for future use VE-5082@CA-ABThe Great Trail of Canada (the Canadian Trailway) National Scenic Trail, within Alberta, Canada 8P-0012Chancery Lane Swamp National Park VK-0556Pieman River State Reserve K-4562@US-CAPacific Crest Trail, within California, USA Additional Notes on POTARef: A browsable and searchable list of all park references is available. A complete CSV file is available (generated nightly). For more information, visit the Parks on the Air documentation website.
	POTARefList                               DataType = "POTARefList"                               // POTARefList = a comma-delimited list of one or more POTARef items.
	SecondarySubdivisionList                  DataType = "SecondarySubdivisionList"                  // SecondarySubdivisionList = a colon-delimited list of two or more members of the Secondary_Administrative_Subdivision enumeration. E.g.: MA,Franklin:MA,Hampshire
	SecondaryAdministrativeSubdivisionListAlt DataType = "SecondaryAdministrativeSubdivisionListAlt" // SecondaryAdministrativeSubdivisionListAlt = a semicolon (;) delimited, unordered list of one or more members of a Secondary_Administrative_Subdivision_Alt enumeration in the form: enumeration-name:enumeration-code Where there is more than one locality represented by the enumeration-code, they are separated by slash (/) characters. Only one of each enumeration-name valid for the DXCC entity concerned can appear in the list. Examples: <CNTY_ALT:28>NZ_Regions:Hawkes Bay/Wairoa <MY_CNTY_ALT:52>NZ_Islands:North Island;NZ_Regions:Hawkes Bay/Wairoa The first example shows the enumeration-name NZ_Regions with the region Hawkes Bay and the district Wairoa. For the purposes of illustration, the second example includes a non-existent subdivision with two available enumeration-codes, NZ_Islands:North Island and NZ_Islands:South Island. The example shows: the enumeration-name NZ_Islands with the island North Island the enumeration-name NZ_Regions with the region Hawkes Bay and the district Wairoa
	SOTARef                                   DataType = "SOTARef"                                   // SOTARef = a sequence of Characters representing an International SOTA Reference. The sequence comprises: an ITU prefix if applicable, a SOTA subdivision a / Character a SOTA Reference Number Examples: W2/WE-003 G/LD-003
	WWFFRef                                   DataType = "WWFFRef"                                   // WWFFRef = a sequence of case-insensitive Characters representing an International WWFF (World Wildlife Flora & Fauna) reference in the form xxFF-nnnn comprising 8 to 11 characters where: xx is the WWFF national program and is 1 to 4 characters in length. FF- is two F characters followed by a dash character. nnnn represents the unique number within the national program and is 4 characters in length with leading zeros. Examples: KFF-4655 3DAFF-0002
)
