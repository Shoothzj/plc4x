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
type BACnetServiceAckReadProperty struct {
	*BACnetServiceAck
	ObjectIdentifier   *BACnetContextTagObjectIdentifier
	PropertyIdentifier *BACnetContextTagPropertyIdentifier
	ArrayIndex         *BACnetContextTagUnsignedInteger
	Values             *BACnetConstructedData
}

// The corresponding interface
type IBACnetServiceAckReadProperty interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetServiceAckReadProperty) ServiceChoice() uint8 {
	return 0x0C
}

func (m *BACnetServiceAckReadProperty) InitializeParent(parent *BACnetServiceAck) {}

func NewBACnetServiceAckReadProperty(objectIdentifier *BACnetContextTagObjectIdentifier, propertyIdentifier *BACnetContextTagPropertyIdentifier, arrayIndex *BACnetContextTagUnsignedInteger, values *BACnetConstructedData) *BACnetServiceAck {
	child := &BACnetServiceAckReadProperty{
		ObjectIdentifier:   objectIdentifier,
		PropertyIdentifier: propertyIdentifier,
		ArrayIndex:         arrayIndex,
		Values:             values,
		BACnetServiceAck:   NewBACnetServiceAck(),
	}
	child.Child = child
	return child.BACnetServiceAck
}

func CastBACnetServiceAckReadProperty(structType interface{}) *BACnetServiceAckReadProperty {
	castFunc := func(typ interface{}) *BACnetServiceAckReadProperty {
		if casted, ok := typ.(BACnetServiceAckReadProperty); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetServiceAckReadProperty); ok {
			return casted
		}
		if casted, ok := typ.(BACnetServiceAck); ok {
			return CastBACnetServiceAckReadProperty(casted.Child)
		}
		if casted, ok := typ.(*BACnetServiceAck); ok {
			return CastBACnetServiceAckReadProperty(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetServiceAckReadProperty) GetTypeName() string {
	return "BACnetServiceAckReadProperty"
}

func (m *BACnetServiceAckReadProperty) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetServiceAckReadProperty) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.LengthInBits()

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.LengthInBits()

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += (*m.ArrayIndex).LengthInBits()
	}

	// Optional Field (values)
	if m.Values != nil {
		lengthInBits += (*m.Values).LengthInBits()
	}

	return lengthInBits
}

func (m *BACnetServiceAckReadProperty) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetServiceAckReadPropertyParse(readBuffer utils.ReadBuffer) (*BACnetServiceAck, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckReadProperty"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(0), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field")
	}
	objectIdentifier := CastBACnetContextTagObjectIdentifier(_objectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_BACNET_PROPERTY_IDENTIFIER)
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field")
	}
	propertyIdentifier := CastBACnetContextTagPropertyIdentifier(_propertyIdentifier)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (arrayIndex) (Can be skipped, if a given expression evaluates to false)
	var arrayIndex *BACnetContextTagUnsignedInteger = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("arrayIndex"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(2), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case _err != nil && _err != utils.ParseAssertError:
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field")
		case _err == utils.ParseAssertError:
			readBuffer.Reset(currentPos)
		default:
			arrayIndex = CastBACnetContextTagUnsignedInteger(_val)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Optional Field (values) (Can be skipped, if a given expression evaluates to false)
	var values *BACnetConstructedData = nil
	{
		currentPos := readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("values"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetConstructedDataParse(readBuffer, uint8(3), objectIdentifier.ObjectType, propertyIdentifier)
		switch {
		case _err != nil && _err != utils.ParseAssertError:
			return nil, errors.Wrap(_err, "Error parsing 'values' field")
		case _err == utils.ParseAssertError:
			readBuffer.Reset(currentPos)
		default:
			values = CastBACnetConstructedData(_val)
			if closeErr := readBuffer.CloseContext("values"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckReadProperty"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckReadProperty{
		ObjectIdentifier:   CastBACnetContextTagObjectIdentifier(objectIdentifier),
		PropertyIdentifier: CastBACnetContextTagPropertyIdentifier(propertyIdentifier),
		ArrayIndex:         CastBACnetContextTagUnsignedInteger(arrayIndex),
		Values:             CastBACnetConstructedData(values),
		BACnetServiceAck:   &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child.BACnetServiceAck, nil
}

func (m *BACnetServiceAckReadProperty) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckReadProperty"); pushErr != nil {
			return pushErr
		}

		// Simple Field (objectIdentifier)
		if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
			return pushErr
		}
		_objectIdentifierErr := m.ObjectIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
			return popErr
		}
		if _objectIdentifierErr != nil {
			return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
		}

		// Simple Field (propertyIdentifier)
		if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
			return pushErr
		}
		_propertyIdentifierErr := m.PropertyIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
			return popErr
		}
		if _propertyIdentifierErr != nil {
			return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
		}

		// Optional Field (arrayIndex) (Can be skipped, if the value is null)
		var arrayIndex *BACnetContextTagUnsignedInteger = nil
		if m.ArrayIndex != nil {
			if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
				return pushErr
			}
			arrayIndex = m.ArrayIndex
			_arrayIndexErr := arrayIndex.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
				return popErr
			}
			if _arrayIndexErr != nil {
				return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
			}
		}

		// Optional Field (values) (Can be skipped, if the value is null)
		var values *BACnetConstructedData = nil
		if m.Values != nil {
			if pushErr := writeBuffer.PushContext("values"); pushErr != nil {
				return pushErr
			}
			values = m.Values
			_valuesErr := values.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("values"); popErr != nil {
				return popErr
			}
			if _valuesErr != nil {
				return errors.Wrap(_valuesErr, "Error serializing 'values' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckReadProperty"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckReadProperty) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
