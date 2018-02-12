package opsutilsgo

import (
	"fmt"
	"sort"
)

type IPAcl struct {
	StartNum, StopNum uint32
	StartIP, StopIP   string
	Info              []string
}

type IPAclTable []IPAcl

func (m IPAclTable) Len() int {
	return len(m)
}

func (m IPAclTable) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m IPAclTable) Less(i, j int) bool {
	if m[i].StartNum == m[j].StartNum {
		return m[i].StopNum < m[j].StopNum
	}
	return m[i].StartNum < m[j].StartNum
}

func (m *IPAclTable) InitByCsv(filepath string) error {
	content, err := CsvReader(filepath)
	if err != nil {
		return err
	}
	lines := content[1:]
	if err := m.Init(lines); err != nil {
		return err
	}
	return nil
}

func (m *IPAclTable) Init(lines [][]string) error {
	for _, line := range lines {
		ipStartNum, err1 := IPStr2Int(line[0])
		ipStopNum, err2 := IPStr2Int(line[1])
		if (err1 != nil) && (err2 != nil) {
			continue
		}
		*m = append(*m, IPAcl{ipStartNum, ipStopNum, line[0], line[1], line})
	}
	sort.Sort(*m)
	return nil
}

func (m IPAclTable) IPLookup(ip string) (bool, int, error) {
	ipNum, err := IPStr2Int(ip)
	if err != nil {
		return false, -1, err
	}
	b := sort.Search(len(m), func(i int) bool { return m[i].StopNum >= ipNum })
	if b != len(m) && m[b].StartNum <= ipNum {
		return true, b, nil
	} else {
		return false, -1, nil
	}
}

func TestIPAclTable() {
	var ipAT IPAclTable
	ipAT.InitByCsv(`D:\workspace\go\workspace\src\github.com\fourth04\demo\ip.merge.csv`)
	ips := []string{
		"198.2.215.125",
		"213.155.156.189",
		"213.155.156.188",
		"45.40.164.134",
		"123412342134234234434234",
	}
	for _, ip := range ips {
		result, index, _ := ipAT.IPLookup(ip)
		fmt.Println(result, index)
	}
}
