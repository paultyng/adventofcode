package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPacketUnmarshalJSON(t *testing.T) {
	for _, tc := range []struct {
		json     string
		expected packet
	}{
		{"[1,1,3,1,1]", packet{packetData{Int: 1}, packetData{Int: 1}, packetData{Int: 3}, packetData{Int: 1}, packetData{Int: 1}}},
		{"[[1],[2,3,4]]", packet{packetData{List: []packetData{{Int: 1}}}, packetData{List: []packetData{{Int: 2}, {Int: 3}, {Int: 4}}}}},
		{"[[[]]]", packet{packetData{List: []packetData{{List: []packetData{}}}}}},

		// divider packets
		{"[[2]]", packet{packetData{List: []packetData{{Int: 2}}}}},
		{"[[6]]", packet{packetData{List: []packetData{{Int: 6}}}}},
	} {
		t.Run(tc.json, func(t *testing.T) {
			var p packet
			err := json.Unmarshal([]byte(tc.json), &p)
			if err != nil {
				t.Fatalf("unable to unmarshal packet: %v", err)
			}
			if !reflect.DeepEqual(p, tc.expected) {
				t.Fatalf("unexpected packet: got %v, want %v", p, tc.expected)
			}
		})
	}
}
