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
type BACnetUnconfirmedServiceRequestIAm struct {
	*BACnetUnconfirmedServiceRequest
	DeviceIdentifier                *BACnetApplicationTagObjectIdentifier
	MaximumApduLengthAcceptedLength *BACnetApplicationTagUnsignedInteger
	SegmentationSupported           *BACnetSegmentation
	VendorId                        *BACnetApplicationTagUnsignedInteger
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestIAm interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetUnconfirmedServiceRequestIAm) ServiceChoice() uint8 {
	return 0x00
}

func (m *BACnetUnconfirmedServiceRequestIAm) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func NewBACnetUnconfirmedServiceRequestIAm(deviceIdentifier *BACnetApplicationTagObjectIdentifier, maximumApduLengthAcceptedLength *BACnetApplicationTagUnsignedInteger, segmentationSupported *BACnetSegmentation, vendorId *BACnetApplicationTagUnsignedInteger) *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestIAm{
		DeviceIdentifier:                deviceIdentifier,
		MaximumApduLengthAcceptedLength: maximumApduLengthAcceptedLength,
		SegmentationSupported:           segmentationSupported,
		VendorId:                        vendorId,
		BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(),
	}
	child.Child = child
	return child.BACnetUnconfirmedServiceRequest
}

func CastBACnetUnconfirmedServiceRequestIAm(structType interface{}) *BACnetUnconfirmedServiceRequestIAm {
	castFunc := func(typ interface{}) *BACnetUnconfirmedServiceRequestIAm {
		if casted, ok := typ.(BACnetUnconfirmedServiceRequestIAm); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequestIAm); ok {
			return casted
		}
		if casted, ok := typ.(BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestIAm(casted.Child)
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestIAm(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetUnconfirmedServiceRequestIAm) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestIAm"
}

func (m *BACnetUnconfirmedServiceRequestIAm) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestIAm) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (deviceIdentifier)
	lengthInBits += m.DeviceIdentifier.LengthInBits()

	// Simple field (maximumApduLengthAcceptedLength)
	lengthInBits += m.MaximumApduLengthAcceptedLength.LengthInBits()

	// Simple field (segmentationSupported)
	lengthInBits += m.SegmentationSupported.LengthInBits()

	// Simple field (vendorId)
	lengthInBits += m.VendorId.LengthInBits()

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestIAm) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestIAmParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetUnconfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestIAm"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (deviceIdentifier)
	if pullErr := readBuffer.PullContext("deviceIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_deviceIdentifier, _deviceIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _deviceIdentifierErr != nil {
		return nil, errors.Wrap(_deviceIdentifierErr, "Error parsing 'deviceIdentifier' field")
	}
	deviceIdentifier := CastBACnetApplicationTagObjectIdentifier(_deviceIdentifier)
	if closeErr := readBuffer.CloseContext("deviceIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (maximumApduLengthAcceptedLength)
	if pullErr := readBuffer.PullContext("maximumApduLengthAcceptedLength"); pullErr != nil {
		return nil, pullErr
	}
	_maximumApduLengthAcceptedLength, _maximumApduLengthAcceptedLengthErr := BACnetApplicationTagParse(readBuffer)
	if _maximumApduLengthAcceptedLengthErr != nil {
		return nil, errors.Wrap(_maximumApduLengthAcceptedLengthErr, "Error parsing 'maximumApduLengthAcceptedLength' field")
	}
	maximumApduLengthAcceptedLength := CastBACnetApplicationTagUnsignedInteger(_maximumApduLengthAcceptedLength)
	if closeErr := readBuffer.CloseContext("maximumApduLengthAcceptedLength"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (segmentationSupported)
	if pullErr := readBuffer.PullContext("segmentationSupported"); pullErr != nil {
		return nil, pullErr
	}
	_segmentationSupported, _segmentationSupportedErr := BACnetSegmentationParse(readBuffer)
	if _segmentationSupportedErr != nil {
		return nil, errors.Wrap(_segmentationSupportedErr, "Error parsing 'segmentationSupported' field")
	}
	segmentationSupported := CastBACnetSegmentation(_segmentationSupported)
	if closeErr := readBuffer.CloseContext("segmentationSupported"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (vendorId)
	if pullErr := readBuffer.PullContext("vendorId"); pullErr != nil {
		return nil, pullErr
	}
	_vendorId, _vendorIdErr := BACnetApplicationTagParse(readBuffer)
	if _vendorIdErr != nil {
		return nil, errors.Wrap(_vendorIdErr, "Error parsing 'vendorId' field")
	}
	vendorId := CastBACnetApplicationTagUnsignedInteger(_vendorId)
	if closeErr := readBuffer.CloseContext("vendorId"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestIAm"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestIAm{
		DeviceIdentifier:                CastBACnetApplicationTagObjectIdentifier(deviceIdentifier),
		MaximumApduLengthAcceptedLength: CastBACnetApplicationTagUnsignedInteger(maximumApduLengthAcceptedLength),
		SegmentationSupported:           CastBACnetSegmentation(segmentationSupported),
		VendorId:                        CastBACnetApplicationTagUnsignedInteger(vendorId),
		BACnetUnconfirmedServiceRequest: &BACnetUnconfirmedServiceRequest{},
	}
	_child.BACnetUnconfirmedServiceRequest.Child = _child
	return _child.BACnetUnconfirmedServiceRequest, nil
}

func (m *BACnetUnconfirmedServiceRequestIAm) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestIAm"); pushErr != nil {
			return pushErr
		}

		// Simple Field (deviceIdentifier)
		if pushErr := writeBuffer.PushContext("deviceIdentifier"); pushErr != nil {
			return pushErr
		}
		_deviceIdentifierErr := m.DeviceIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("deviceIdentifier"); popErr != nil {
			return popErr
		}
		if _deviceIdentifierErr != nil {
			return errors.Wrap(_deviceIdentifierErr, "Error serializing 'deviceIdentifier' field")
		}

		// Simple Field (maximumApduLengthAcceptedLength)
		if pushErr := writeBuffer.PushContext("maximumApduLengthAcceptedLength"); pushErr != nil {
			return pushErr
		}
		_maximumApduLengthAcceptedLengthErr := m.MaximumApduLengthAcceptedLength.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("maximumApduLengthAcceptedLength"); popErr != nil {
			return popErr
		}
		if _maximumApduLengthAcceptedLengthErr != nil {
			return errors.Wrap(_maximumApduLengthAcceptedLengthErr, "Error serializing 'maximumApduLengthAcceptedLength' field")
		}

		// Simple Field (segmentationSupported)
		if pushErr := writeBuffer.PushContext("segmentationSupported"); pushErr != nil {
			return pushErr
		}
		_segmentationSupportedErr := m.SegmentationSupported.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("segmentationSupported"); popErr != nil {
			return popErr
		}
		if _segmentationSupportedErr != nil {
			return errors.Wrap(_segmentationSupportedErr, "Error serializing 'segmentationSupported' field")
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

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestIAm"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestIAm) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
