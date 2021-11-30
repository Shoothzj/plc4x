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
type BACnetTagApplicationUnsignedInteger struct {
	ValueUint8  *uint8
	ValueUint16 *uint16
	ValueUint32 *uint32
	ValueUint64 *uint64
	IsUint8     bool
	IsUint16    bool
	IsUint32    bool
	IsUint64    bool
	Parent      *BACnetTag
}

// The corresponding interface
type IBACnetTagApplicationUnsignedInteger interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetTagApplicationUnsignedInteger) TagClass() TagClass {
	return TagClass_APPLICATION_TAGS
}

func (m *BACnetTagApplicationUnsignedInteger) InitializeParent(parent *BACnetTag, tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32, actualTagNumber uint8, isPrimitiveAndNotBoolean bool, actualLength uint32) {
	m.Parent.TagNumber = tagNumber
	m.Parent.LengthValueType = lengthValueType
	m.Parent.ExtTagNumber = extTagNumber
	m.Parent.ExtLength = extLength
	m.Parent.ExtExtLength = extExtLength
	m.Parent.ExtExtExtLength = extExtExtLength
}

func NewBACnetTagApplicationUnsignedInteger(valueUint8 *uint8, valueUint16 *uint16, valueUint32 *uint32, valueUint64 *uint64, tagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetTag {
	child := &BACnetTagApplicationUnsignedInteger{
		ValueUint8:  valueUint8,
		ValueUint16: valueUint16,
		ValueUint32: valueUint32,
		ValueUint64: valueUint64,
		Parent:      NewBACnetTag(tagNumber, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetTagApplicationUnsignedInteger(structType interface{}) *BACnetTagApplicationUnsignedInteger {
	castFunc := func(typ interface{}) *BACnetTagApplicationUnsignedInteger {
		if casted, ok := typ.(BACnetTagApplicationUnsignedInteger); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTagApplicationUnsignedInteger); ok {
			return casted
		}
		if casted, ok := typ.(BACnetTag); ok {
			return CastBACnetTagApplicationUnsignedInteger(casted.Child)
		}
		if casted, ok := typ.(*BACnetTag); ok {
			return CastBACnetTagApplicationUnsignedInteger(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTagApplicationUnsignedInteger) GetTypeName() string {
	return "BACnetTagApplicationUnsignedInteger"
}

func (m *BACnetTagApplicationUnsignedInteger) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetTagApplicationUnsignedInteger) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint8)
	if m.ValueUint8 != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint16)
	if m.ValueUint16 != nil {
		lengthInBits += 16
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint32)
	if m.ValueUint32 != nil {
		lengthInBits += 32
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (valueUint64)
	if m.ValueUint64 != nil {
		lengthInBits += 64
	}

	return lengthInBits
}

func (m *BACnetTagApplicationUnsignedInteger) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetTagApplicationUnsignedIntegerParse(readBuffer utils.ReadBuffer, actualLength uint32) (*BACnetTag, error) {
	if pullErr := readBuffer.PullContext("BACnetTagApplicationUnsignedInteger"); pullErr != nil {
		return nil, pullErr
	}

	// Virtual field
	_isUint8 := bool((actualLength) == (1))
	isUint8 := bool(_isUint8)

	// Optional Field (valueUint8) (Can be skipped, if a given expression evaluates to false)
	var valueUint8 *uint8 = nil
	if isUint8 {
		_val, _err := readBuffer.ReadUint8("valueUint8", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint8' field")
		}
		valueUint8 = &_val
	}

	// Virtual field
	_isUint16 := bool((actualLength) == (2))
	isUint16 := bool(_isUint16)

	// Optional Field (valueUint16) (Can be skipped, if a given expression evaluates to false)
	var valueUint16 *uint16 = nil
	if isUint16 {
		_val, _err := readBuffer.ReadUint16("valueUint16", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint16' field")
		}
		valueUint16 = &_val
	}

	// Virtual field
	_isUint32 := bool((actualLength) == (3))
	isUint32 := bool(_isUint32)

	// Optional Field (valueUint32) (Can be skipped, if a given expression evaluates to false)
	var valueUint32 *uint32 = nil
	if isUint32 {
		_val, _err := readBuffer.ReadUint32("valueUint32", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint32' field")
		}
		valueUint32 = &_val
	}

	// Virtual field
	_isUint64 := bool((actualLength) == (4))
	isUint64 := bool(_isUint64)

	// Optional Field (valueUint64) (Can be skipped, if a given expression evaluates to false)
	var valueUint64 *uint64 = nil
	if isUint64 {
		_val, _err := readBuffer.ReadUint64("valueUint64", 64)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'valueUint64' field")
		}
		valueUint64 = &_val
	}

	if closeErr := readBuffer.CloseContext("BACnetTagApplicationUnsignedInteger"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetTagApplicationUnsignedInteger{
		ValueUint8:  valueUint8,
		ValueUint16: valueUint16,
		ValueUint32: valueUint32,
		ValueUint64: valueUint64,
		IsUint8:     isUint8,
		IsUint16:    isUint16,
		IsUint32:    isUint32,
		IsUint64:    isUint64,
		Parent:      &BACnetTag{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetTagApplicationUnsignedInteger) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTagApplicationUnsignedInteger"); pushErr != nil {
			return pushErr
		}

		// Optional Field (valueUint8) (Can be skipped, if the value is null)
		var valueUint8 *uint8 = nil
		if m.ValueUint8 != nil {
			valueUint8 = m.ValueUint8
			_valueUint8Err := writeBuffer.WriteUint8("valueUint8", 8, *(valueUint8))
			if _valueUint8Err != nil {
				return errors.Wrap(_valueUint8Err, "Error serializing 'valueUint8' field")
			}
		}

		// Optional Field (valueUint16) (Can be skipped, if the value is null)
		var valueUint16 *uint16 = nil
		if m.ValueUint16 != nil {
			valueUint16 = m.ValueUint16
			_valueUint16Err := writeBuffer.WriteUint16("valueUint16", 16, *(valueUint16))
			if _valueUint16Err != nil {
				return errors.Wrap(_valueUint16Err, "Error serializing 'valueUint16' field")
			}
		}

		// Optional Field (valueUint32) (Can be skipped, if the value is null)
		var valueUint32 *uint32 = nil
		if m.ValueUint32 != nil {
			valueUint32 = m.ValueUint32
			_valueUint32Err := writeBuffer.WriteUint32("valueUint32", 32, *(valueUint32))
			if _valueUint32Err != nil {
				return errors.Wrap(_valueUint32Err, "Error serializing 'valueUint32' field")
			}
		}

		// Optional Field (valueUint64) (Can be skipped, if the value is null)
		var valueUint64 *uint64 = nil
		if m.ValueUint64 != nil {
			valueUint64 = m.ValueUint64
			_valueUint64Err := writeBuffer.WriteUint64("valueUint64", 64, *(valueUint64))
			if _valueUint64Err != nil {
				return errors.Wrap(_valueUint64Err, "Error serializing 'valueUint64' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetTagApplicationUnsignedInteger"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetTagApplicationUnsignedInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
