package id

import "testing"

func TestNewUniqueIDGenerator(t *testing.T) {
	tests := []struct {
		name              string
		dataCenterID      int
		nodeID            int
		epoch             int64
		expectedGenerator *UniqueIDGenerator
		expectError       bool
	}{
		{"Negative_DataCenter_ID", -1, 0, 0, nil, true},
		{"Too_Large_DataCenter_ID", 32, 0, 0, nil, true},
		{"Negative_Node_ID", 0, -1, 0, nil, true},
		{"Too_Large_Node_ID", 0, 32, 0, nil, true},
		{"Negative_Epoch", 0, 0, -1, nil, true},
		{"Default", 0, 0, 0, &UniqueIDGenerator{}, false},
		{"Created", 11, 17, 946684800000, &UniqueIDGenerator{config: config{11, 17, 946684800000}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator, err := NewUniqueIDGenerator(tt.dataCenterID, tt.nodeID, tt.epoch)
			if (err != nil) != tt.expectError {
				t.Errorf("UniqueIDGenerator(%v, %v, %v) returned unexpected error status: got error %v, expected error status %v", tt.dataCenterID, tt.nodeID, tt.epoch, err, tt.expectError)
			}
			if !tt.expectError && (generator == nil || *generator != *tt.expectedGenerator) {
				t.Errorf("UniqueIDGenerator(%v, %v, %v) = %v, expected %v", tt.dataCenterID, tt.nodeID, tt.epoch, generator, tt.expectedGenerator)
			}
		})
	}
}
