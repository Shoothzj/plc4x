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
type BACnetNotificationParametersExtended struct {
	*BACnetNotificationParameters
	InnerOpeningTag   *BACnetOpeningTag
	VendorId          *BACnetContextTagUnsignedInteger
	ExtendedEventType *BACnetContextTagUnsignedInteger
	Parameters        *BACnetNotificationParametersExtendedParameters
	InnerClosingTag   *BACnetClosingTag
}

// The corresponding interface
type IBACnetNotificationParametersExtended interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetNotificationParametersExtended) PeekedTagNumber() uint8 {
	return uint8(9)
}

func (m *BACnetNotificationParametersExtended) InitializeParent(parent *BACnetNotificationParameters, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, peekedTagNumber uint8) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func NewBACnetNotificationParametersExtended(innerOpeningTag *BACnetOpeningTag, vendorId *BACnetContextTagUnsignedInteger, extendedEventType *BACnetContextTagUnsignedInteger, parameters *BACnetNotificationParametersExtendedParameters, innerClosingTag *BACnetClosingTag, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, peekedTagNumber uint8) *BACnetNotificationParameters {
	child := &BACnetNotificationParametersExtended{
		InnerOpeningTag:              innerOpeningTag,
		VendorId:                     vendorId,
		ExtendedEventType:            extendedEventType,
		Parameters:                   parameters,
		InnerClosingTag:              innerClosingTag,
		BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, peekedTagNumber),
	}
	child.Child = child
	return child.BACnetNotificationParameters
}

func CastBACnetNotificationParametersExtended(structType interface{}) *BACnetNotificationParametersExtended {
	castFunc := func(typ interface{}) *BACnetNotificationParametersExtended {
		if casted, ok := typ.(BACnetNotificationParametersExtended); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetNotificationParametersExtended); ok {
			return casted
		}
		if casted, ok := typ.(BACnetNotificationParameters); ok {
			return CastBACnetNotificationParametersExtended(casted.Child)
		}
		if casted, ok := typ.(*BACnetNotificationParameters); ok {
			return CastBACnetNotificationParametersExtended(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetNotificationParametersExtended) GetTypeName() string {
	return "BACnetNotificationParametersExtended"
}

func (m *BACnetNotificationParametersExtended) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersExtended) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.LengthInBits()

	// Simple field (vendorId)
	lengthInBits += m.VendorId.LengthInBits()

	// Simple field (extendedEventType)
	lengthInBits += m.ExtendedEventType.LengthInBits()

	// Simple field (parameters)
	lengthInBits += m.Parameters.LengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.LengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersExtended) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetNotificationParametersExtendedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, peekedTagNumber uint8) (*BACnetNotificationParameters, error) {
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersExtended"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, pullErr
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetContextTagParse(readBuffer, peekedTagNumber, BACnetDataType_OPENING_TAG)
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field")
	}
	innerOpeningTag := CastBACnetOpeningTag(_innerOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (vendorId)
	if pullErr := readBuffer.PullContext("vendorId"); pullErr != nil {
		return nil, pullErr
	}
	_vendorId, _vendorIdErr := BACnetContextTagParse(readBuffer, uint8(0), BACnetDataType_UNSIGNED_INTEGER)
	if _vendorIdErr != nil {
		return nil, errors.Wrap(_vendorIdErr, "Error parsing 'vendorId' field")
	}
	vendorId := CastBACnetContextTagUnsignedInteger(_vendorId)
	if closeErr := readBuffer.CloseContext("vendorId"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (extendedEventType)
	if pullErr := readBuffer.PullContext("extendedEventType"); pullErr != nil {
		return nil, pullErr
	}
	_extendedEventType, _extendedEventTypeErr := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_UNSIGNED_INTEGER)
	if _extendedEventTypeErr != nil {
		return nil, errors.Wrap(_extendedEventTypeErr, "Error parsing 'extendedEventType' field")
	}
	extendedEventType := CastBACnetContextTagUnsignedInteger(_extendedEventType)
	if closeErr := readBuffer.CloseContext("extendedEventType"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (parameters)
	if pullErr := readBuffer.PullContext("parameters"); pullErr != nil {
		return nil, pullErr
	}
	_parameters, _parametersErr := BACnetNotificationParametersExtendedParametersParse(readBuffer, uint8(2))
	if _parametersErr != nil {
		return nil, errors.Wrap(_parametersErr, "Error parsing 'parameters' field")
	}
	parameters := CastBACnetNotificationParametersExtendedParameters(_parameters)
	if closeErr := readBuffer.CloseContext("parameters"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, pullErr
	}
	_innerClosingTag, _innerClosingTagErr := BACnetContextTagParse(readBuffer, peekedTagNumber, BACnetDataType_CLOSING_TAG)
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field")
	}
	innerClosingTag := CastBACnetClosingTag(_innerClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersExtended"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersExtended{
		InnerOpeningTag:              CastBACnetOpeningTag(innerOpeningTag),
		VendorId:                     CastBACnetContextTagUnsignedInteger(vendorId),
		ExtendedEventType:            CastBACnetContextTagUnsignedInteger(extendedEventType),
		Parameters:                   CastBACnetNotificationParametersExtendedParameters(parameters),
		InnerClosingTag:              CastBACnetClosingTag(innerClosingTag),
		BACnetNotificationParameters: &BACnetNotificationParameters{},
	}
	_child.BACnetNotificationParameters.Child = _child
	return _child.BACnetNotificationParameters, nil
}

func (m *BACnetNotificationParametersExtended) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersExtended"); pushErr != nil {
			return pushErr
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return pushErr
		}
		_innerOpeningTagErr := m.InnerOpeningTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return popErr
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (vendorId)
		if pushErr := writeBuffer.PushContext("vendorId"); pushErr != nil {
			return pushErr
		}
		_vendorIdErr := m.VendorId.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("vendorId"); popErr != nil {
			return popErr
		}
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		// Simple Field (extendedEventType)
		if pushErr := writeBuffer.PushContext("extendedEventType"); pushErr != nil {
			return pushErr
		}
		_extendedEventTypeErr := m.ExtendedEventType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("extendedEventType"); popErr != nil {
			return popErr
		}
		if _extendedEventTypeErr != nil {
			return errors.Wrap(_extendedEventTypeErr, "Error serializing 'extendedEventType' field")
		}

		// Simple Field (parameters)
		if pushErr := writeBuffer.PushContext("parameters"); pushErr != nil {
			return pushErr
		}
		_parametersErr := m.Parameters.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("parameters"); popErr != nil {
			return popErr
		}
		if _parametersErr != nil {
			return errors.Wrap(_parametersErr, "Error serializing 'parameters' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return pushErr
		}
		_innerClosingTagErr := m.InnerClosingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return popErr
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersExtended"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
