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
type BACnetNotificationParametersComplexEventType struct {
	*BACnetNotificationParameters
	ListOfValues *BACnetPropertyValues
}

// The corresponding interface
type IBACnetNotificationParametersComplexEventType interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetNotificationParametersComplexEventType) PeekedTagNumber() uint8 {
	return uint8(6)
}

func (m *BACnetNotificationParametersComplexEventType) InitializeParent(parent *BACnetNotificationParameters, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, peekedTagNumber uint8) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func NewBACnetNotificationParametersComplexEventType(listOfValues *BACnetPropertyValues, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, peekedTagNumber uint8) *BACnetNotificationParameters {
	child := &BACnetNotificationParametersComplexEventType{
		ListOfValues:                 listOfValues,
		BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, peekedTagNumber),
	}
	child.Child = child
	return child.BACnetNotificationParameters
}

func CastBACnetNotificationParametersComplexEventType(structType interface{}) *BACnetNotificationParametersComplexEventType {
	castFunc := func(typ interface{}) *BACnetNotificationParametersComplexEventType {
		if casted, ok := typ.(BACnetNotificationParametersComplexEventType); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetNotificationParametersComplexEventType); ok {
			return casted
		}
		if casted, ok := typ.(BACnetNotificationParameters); ok {
			return CastBACnetNotificationParametersComplexEventType(casted.Child)
		}
		if casted, ok := typ.(*BACnetNotificationParameters); ok {
			return CastBACnetNotificationParametersComplexEventType(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetNotificationParametersComplexEventType) GetTypeName() string {
	return "BACnetNotificationParametersComplexEventType"
}

func (m *BACnetNotificationParametersComplexEventType) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersComplexEventType) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (listOfValues)
	lengthInBits += m.ListOfValues.LengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersComplexEventType) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetNotificationParametersComplexEventTypeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, peekedTagNumber uint8) (*BACnetNotificationParameters, error) {
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersComplexEventType"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (listOfValues)
	if pullErr := readBuffer.PullContext("listOfValues"); pullErr != nil {
		return nil, pullErr
	}
	_listOfValues, _listOfValuesErr := BACnetPropertyValuesParse(readBuffer, peekedTagNumber, objectType)
	if _listOfValuesErr != nil {
		return nil, errors.Wrap(_listOfValuesErr, "Error parsing 'listOfValues' field")
	}
	listOfValues := CastBACnetPropertyValues(_listOfValues)
	if closeErr := readBuffer.CloseContext("listOfValues"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersComplexEventType"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersComplexEventType{
		ListOfValues:                 CastBACnetPropertyValues(listOfValues),
		BACnetNotificationParameters: &BACnetNotificationParameters{},
	}
	_child.BACnetNotificationParameters.Child = _child
	return _child.BACnetNotificationParameters, nil
}

func (m *BACnetNotificationParametersComplexEventType) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersComplexEventType"); pushErr != nil {
			return pushErr
		}

		// Simple Field (listOfValues)
		if pushErr := writeBuffer.PushContext("listOfValues"); pushErr != nil {
			return pushErr
		}
		_listOfValuesErr := m.ListOfValues.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("listOfValues"); popErr != nil {
			return popErr
		}
		if _listOfValuesErr != nil {
			return errors.Wrap(_listOfValuesErr, "Error serializing 'listOfValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersComplexEventType"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersComplexEventType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
