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
type BACnetTimeStamp struct {
	OpeningTag      *BACnetOpeningTag
	PeekedTagHeader *BACnetTagHeader
	ClosingTag      *BACnetClosingTag
	PeekedTagNumber uint8
	Child           IBACnetTimeStampChild
}

// The corresponding interface
type IBACnetTimeStamp interface {
	PeekedTagNumber() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetTimeStampParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetTimeStamp, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetTimeStampChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetTimeStamp, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, peekedTagNumber uint8)
	GetTypeName() string
	IBACnetTimeStamp
}

func NewBACnetTimeStamp(openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, peekedTagNumber uint8) *BACnetTimeStamp {
	return &BACnetTimeStamp{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, PeekedTagNumber: peekedTagNumber}
}

func CastBACnetTimeStamp(structType interface{}) *BACnetTimeStamp {
	castFunc := func(typ interface{}) *BACnetTimeStamp {
		if casted, ok := typ.(BACnetTimeStamp); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTimeStamp); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTimeStamp) GetTypeName() string {
	return "BACnetTimeStamp"
}

func (m *BACnetTimeStamp) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetTimeStamp) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *BACnetTimeStamp) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.LengthInBits()

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.LengthInBits()

	return lengthInBits
}

func (m *BACnetTimeStamp) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetTimeStampParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetTimeStamp, error) {
	if pullErr := readBuffer.PullContext("BACnetTimeStamp"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetContextTagParse(readBuffer, tagNumber, BACnetDataType_OPENING_TAG)
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Peek Field (peekedTagHeader)
	currentPos := readBuffer.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, pullErr
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.ActualTagNumber
	peekedTagNumber := uint8(_peekedTagNumber)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetTimeStamp
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetTimeStampTime
		_parent, typeSwitchError = BACnetTimeStampTimeParse(readBuffer, tagNumber)
	case peekedTagNumber == uint8(1): // BACnetTimeStampSequence
		_parent, typeSwitchError = BACnetTimeStampSequenceParse(readBuffer, tagNumber)
	case peekedTagNumber == uint8(2): // BACnetTimeStampDateTime
		_parent, typeSwitchError = BACnetTimeStampDateTimeParse(readBuffer, tagNumber)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetContextTagParse(readBuffer, tagNumber, BACnetDataType_CLOSING_TAG)
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetTimeStamp"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, openingTag, peekedTagHeader, closingTag, peekedTagNumber)
	return _parent, nil
}

func (m *BACnetTimeStamp) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetTimeStamp) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetTimeStamp, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetTimeStamp"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.PeekedTagNumber); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTimeStamp"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTimeStamp) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
