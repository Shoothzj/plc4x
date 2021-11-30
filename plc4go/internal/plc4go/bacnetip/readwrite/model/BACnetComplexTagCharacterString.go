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
type BACnetComplexTagCharacterString struct {
	Encoding          BACnetCharacterEncoding
	Value             string
	ActualLengthInBit uint16
	Parent            *BACnetComplexTag
}

// The corresponding interface
type IBACnetComplexTagCharacterString interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetComplexTagCharacterString) DataType() BACnetDataType {
	return BACnetDataType_CHARACTER_STRING
}

func (m *BACnetComplexTagCharacterString) InitializeParent(parent *BACnetComplexTag, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, isPrimitiveAndNotBoolean bool, actualLength uint32) {
	m.Parent.TagNumber = tagNumber
	m.Parent.TagClass = tagClass
	m.Parent.LengthValueType = lengthValueType
	m.Parent.ExtTagNumber = extTagNumber
	m.Parent.ExtLength = extLength
	m.Parent.ExtExtLength = extExtLength
	m.Parent.ExtExtExtLength = extExtExtLength
}

func NewBACnetComplexTagCharacterString(encoding BACnetCharacterEncoding, value string, tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetComplexTag {
	child := &BACnetComplexTagCharacterString{
		Encoding: encoding,
		Value:    value,
		Parent:   NewBACnetComplexTag(tagNumber, tagClass, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetComplexTagCharacterString(structType interface{}) *BACnetComplexTagCharacterString {
	castFunc := func(typ interface{}) *BACnetComplexTagCharacterString {
		if casted, ok := typ.(BACnetComplexTagCharacterString); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetComplexTagCharacterString); ok {
			return casted
		}
		if casted, ok := typ.(BACnetComplexTag); ok {
			return CastBACnetComplexTagCharacterString(casted.Child)
		}
		if casted, ok := typ.(*BACnetComplexTag); ok {
			return CastBACnetComplexTagCharacterString(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetComplexTagCharacterString) GetTypeName() string {
	return "BACnetComplexTagCharacterString"
}

func (m *BACnetComplexTagCharacterString) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetComplexTagCharacterString) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (encoding)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (value)
	lengthInBits += uint16(m.ActualLengthInBit)

	return lengthInBits
}

func (m *BACnetComplexTagCharacterString) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetComplexTagCharacterStringParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, actualLength uint32) (*BACnetComplexTag, error) {
	if pullErr := readBuffer.PullContext("BACnetComplexTagCharacterString"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (encoding)
	if pullErr := readBuffer.PullContext("encoding"); pullErr != nil {
		return nil, pullErr
	}
	encoding, _encodingErr := BACnetCharacterEncodingParse(readBuffer)
	if _encodingErr != nil {
		return nil, errors.Wrap(_encodingErr, "Error parsing 'encoding' field")
	}
	if closeErr := readBuffer.CloseContext("encoding"); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_actualLengthInBit := uint16(uint16(actualLength)*uint16(uint16(8))) - uint16(uint16(8))
	actualLengthInBit := uint16(_actualLengthInBit)

	// Simple Field (value)
	value, _valueErr := readBuffer.ReadString("value", uint32(actualLengthInBit))
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}

	if closeErr := readBuffer.CloseContext("BACnetComplexTagCharacterString"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetComplexTagCharacterString{
		Encoding:          encoding,
		Value:             value,
		ActualLengthInBit: actualLengthInBit,
		Parent:            &BACnetComplexTag{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetComplexTagCharacterString) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetComplexTagCharacterString"); pushErr != nil {
			return pushErr
		}

		// Simple Field (encoding)
		if pushErr := writeBuffer.PushContext("encoding"); pushErr != nil {
			return pushErr
		}
		_encodingErr := m.Encoding.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("encoding"); popErr != nil {
			return popErr
		}
		if _encodingErr != nil {
			return errors.Wrap(_encodingErr, "Error serializing 'encoding' field")
		}

		// Simple Field (value)
		value := string(m.Value)
		_valueErr := writeBuffer.WriteString("value", uint32(m.ActualLengthInBit), "UTF-8", (value))
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("BACnetComplexTagCharacterString"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetComplexTagCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
