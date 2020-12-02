package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	entries := map[int]bool{
		1630: true,
		1801: true,
		1917: true,
		1958: true,
		1953: true,
		1521: true,
		1990: true,
		1959: true,
		1543: true,
		1798: true,
		638:  true,
		1499: true,
		1977: true,
		1433: true,
		1532: true,
		1780: true,
		1559: true,
		1866: true,
		1962: true,
		1999: true,
		1623: true,
		1772: true,
		1730: true,
		1670: true,
		1791: true,
		1947: true,
		1961: true,
		1523: true,
		959:  true,
		1998: true,
		1693: true,
		1490: true,
		1712: true,
		910:  true,
		1635: true,
		1837: true,
		586:  true,
		1590: true,
		1741: true,
		1739: true,
		1660: true,
		1883: true,
		1777: true,
		1734: true,
		1413: true,
		1456: true,
		1511: true,
		1957: true,
		1738: true,
		1685: true,
		1677: true,
		1419: true,
		1566: true,
		1639: true,
		1578: true,
		1922: true,
		1856: true,
		1946: true,
		1965: true,
		1649: true,
		1854: true,
		1610: true,
		1806: true,
		1424: true,
		1616: true,
		218:  true,
		1678: true,
		1992: true,
		1985: true,
		903:  true,
		1626: true,
		1412: true,
		1964: true,
		671:  true,
		1692: true,
		1571: true,
		1690: true,
		1587: true,
		1933: true,
		1367: true,
		1585: true,
		1575: true,
		498:  true,
		1601: true,
		2005: true,
		1711: true,
		1948: true,
		1991: true,
		1580: true,
		1704: true,
		207:  true,
		1560: true,
		1867: true,
		1600: true,
		1594: true,
		1930: true,
		1541: true,
		1832: true,
		1613: true,
		1599: true,
		1757: true,
		71:   true,
		1534: true,
		1940: true,
		1982: true,
		1960: true,
		1530: true,
		1908: true,
		1857: true,
		1410: true,
		1987: true,
		1526: true,
		1546: true,
		2002: true,
		1923: true,
		1972: true,
		1752: true,
		1984: true,
		1754: true,
		1916: true,
		1942: true,
		1980: true,
		1608: true,
		1398: true,
		1438: true,
		1955: true,
		1968: true,
		1799: true,
		1976: true,
		1847: true,
		1775: true,
		1904: true,
		1983: true,
		1945: true,
		1554: true,
		1486: true,
		1527: true,
		1884: true,
		1553: true,
		1736: true,
		1561: true,
		1513: true,
		1695: true,
		1431: true,
		1997: true,
		1405: true,
		1872: true,
		1434: true,
		1679: true,
		1609: true,
		105:  true,
		1582: true,
		1795: true,
		1826: true,
		1886: true,
		1472: true,
		2007: true,
		1617: true,
		1978: true,
		1669: true,
		1764: true,
		1865: true,
		1773: true,
		1993: true,
		1666: true,
		1583: true,
		2009: true,
		1969: true,
		2001: true,
		1659: true,
		1833: true,
		1713: true,
		1893: true,
		2000: true,
		1520: true,
		1652: true,
		1437: true,
		1556: true,
		1633: true,
		1386: true,
		1819: true,
		1973: true,
		1426: true,
		1975: true,
		2010: true,
		1863: true,
		1593: true,
		1996: true,
		1796: true,
		1986: true,
		1995: true,
		657:  true,
		1784: true,
		1644: true,
		1941: true,
		1596: true,
		1849: true,
		1065: true,
		1927: true,
		1525: true,
	}

	sum := 2020

	found, _ := findTwoEntriesWhichSumTo(sum, entries)
	fmt.Println(found)
	fmt.Println(found.first * found.second)

	foundThree, _ := findThreeEntriesWhichSumTo(sum, entries)
	fmt.Println(foundThree)
	fmt.Println(foundThree.first * foundThree.second * foundThree.third)
}

type pair struct {
	first  int
	second int
}

type triplet struct {
	first  int
	second int
	third  int
}

func findTwoEntriesWhichSumTo(sum int, entries map[int]bool) (pair, error) {
	for first := range entries {
		second := sum - first
		_, isSecondFound := entries[second]

		if isSecondFound {
			return pair{first, second}, nil
		}
	}

	return pair{}, errors.New("No pair was found whose sum equals " + strconv.Itoa(sum))
}

func findThreeEntriesWhichSumTo(sum int, entries map[int]bool) (triplet, error) {
	for third := range entries {
		subSum := sum - third

		pair, error := findTwoEntriesWhichSumTo(subSum, entries)

		if error == nil {
			return triplet{pair.first, pair.second, third}, nil
		}
	}

	// Because I don't want to deal with errors right now
	panic("No triplet was found whose sum equals " + strconv.Itoa(sum))
}
