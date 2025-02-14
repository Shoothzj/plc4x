/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetApplicationTagDouble struct {
	*BACnetApplicationTag
	Value float64
}

// The corresponding interface
type IBACnetApplicationTagDouble interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetApplicationTagDouble) TagNumber() uint8 {
	return 0x5
}

func (m *BACnetApplicationTagDouble) InitializeParent(parent *BACnetApplicationTag, header *BACnetTagHeader, tagNumber uint8, actualLength uint32) {
	m.Header = header
}

func NewBACnetApplicationTagDouble(value float64, header *BACnetTagHeader, tagNumber uint8, actualLength uint32) *BACnetApplicationTag {
	child := &BACnetApplicationTagDouble{
		Value:                value,
		BACnetApplicationTag: NewBACnetApplicationTag(header, tagNumber, actualLength),
	}
	child.Child = child
	return child.BACnetApplicationTag
}

func CastBACnetApplicationTagDouble(structType interface{}) *BACnetApplicationTagDouble {
	castFunc := func(typ interface{}) *BACnetApplicationTagDouble {
		if casted, ok := typ.(BACnetApplicationTagDouble); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetApplicationTagDouble); ok {
			return casted
		}
		if casted, ok := typ.(BACnetApplicationTag); ok {
			return CastBACnetApplicationTagDouble(casted.Child)
		}
		if casted, ok := typ.(*BACnetApplicationTag); ok {
			return CastBACnetApplicationTagDouble(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetApplicationTagDouble) GetTypeName() string {
	return "BACnetApplicationTagDouble"
}

func (m *BACnetApplicationTagDouble) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetApplicationTagDouble) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (value)
	lengthInBits += 64

	return lengthInBits
}

func (m *BACnetApplicationTagDouble) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetApplicationTagDoubleParse(readBuffer utils.ReadBuffer) (*BACnetApplicationTag, error) {
	if pullErr := readBuffer.PullContext("BACnetApplicationTagDouble"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadFloat64("value", 64)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagDouble"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetApplicationTagDouble{
		Value:                value,
		BACnetApplicationTag: &BACnetApplicationTag{},
	}
	_child.BACnetApplicationTag.Child = _child
	return _child.BACnetApplicationTag, nil
}

func (m *BACnetApplicationTagDouble) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagDouble"); pushErr != nil {
			return pushErr
		}

		// Simple Field (value)
		value := float64(m.Value)
		_valueErr := writeBuffer.WriteFloat64("value", 64, (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagDouble"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetApplicationTagDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
