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
package org.apache.plc4x.java.s7.utils;

import java.util.HashMap;
import java.util.Map;

/**
 *  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
 *  |15|14|13|12|11|10| 9| 8| 7| 6| 5| 4| 3| 2| 1|
 *  +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
 *  \__________/\__________/\____________________/
 *   Event class     IDs         Event number
 * 
 *   Event Class:
 *      1   Standard OB Events
 *      2   Synchronous errors
 *      3   Asynchronous errors
 *      4   Mode transition
 *      5   Run-time events
 *      6   Communications events
 *      7   Events for fail-safe and fault tolerant systems
 *      8   Standardized diagnostic data on modules
 *      9   Predefined user events
 *    A,B   Freely definable events
 *  C,D,E   Reserved
 *      F   Events for modules other than CPUs (for example, CPs, FMs)
 * 
 *  IDs (Bit)
 *      8   0:Event leaving state, 1:Event entering state
 *      9   1:Entry in diagnostic buffer
 *     10   1:Internal error
 *     11   1:External error
 * 
 * @author cgarcia
 */
public enum S7DiagnosticEventId {

    EVENTID_0x113A((short) 0x113A, "Start request for cyclic interrupt OB with special handling (S7-300 only)"),
    EVENTID_0x1155((short) 0x1155, "Status alarm for PROFIBUS DP"),
    EVENTID_0x1156((short) 0x1156, "Update interrupt for PROFIBUS DP"),
    EVENTID_0x1157((short) 0x1157, "Manufacturer interrupt for PROFIBUS DP"),
    EVENTID_0x1158((short) 0x1158, "Status interrupt for PROFINET IO"),
    EVENTID_0x1159((short) 0x1159, "Update interrupt for PROFINET IO"),
    EVENTID_0x115A((short) 0x115A, "Manufacturer interrupt for PROFINET IO"),
    EVENTID_0x115B((short) 0x115B, "IO: Profile-specific interrupt"),
    EVENTID_0x116A((short) 0x116A, "Technology synchronization interrupt"),
    
    //Event Class 1 - Standard OB Events
    EVENTID_0x1381((short) 0x1381, "Request for manual warm restart"),
    EVENTID_0x1382((short) 0x1382, "Request for automatic warm restart"),
    EVENTID_0x1383((short) 0x1383, "Request for manual hot restart"),
    EVENTID_0x1384((short) 0x1384, "Request for automatic hot restart"),
    EVENTID_0x1385((short) 0x1385, "Request for manual cold restart"),
    EVENTID_0x1386((short) 0x1386, "Request for automatic cold restart"),
    EVENTID_0x1387((short) 0x1387, "Master CPU: request for manual cold restart"),
    EVENTID_0x1388((short) 0x1388, "Master CPU: request for automatic cold restart"),
    EVENTID_0x138A((short) 0x138A, "Master CPU: request for manual warm restart"),
    EVENTID_0x138B((short) 0x138B, "Master CPU: request for automatic warm restart"),
    EVENTID_0x138C((short) 0x138C, "Standby CPU: request for manual hot restart"),
    EVENTID_0x138D((short) 0x138D, "Standby CPU: request for automatic hot restart"),
    
    //Event Class 2 - Synchronous Errors
    EVENTID_0x2521((short) 0x2521, "BCD conversion error"),
    EVENTID_0x2522((short) 0x2522, "Area length error when reading"),
    EVENTID_0x2523((short) 0x2523, "Area length error when writing"),
    EVENTID_0x2524((short) 0x2524, "Area error when reading"),
    EVENTID_0x2525((short) 0x2525, "Area error when writing"),
    EVENTID_0x2526((short) 0x2526, "Timer number error"),
    EVENTID_0x2527((short) 0x2527, "Counter number error"),
    EVENTID_0x2528((short) 0x2528, "Alignment error when reading"),
    EVENTID_0x2529((short) 0x2529, "Alignment error when writing"),
    EVENTID_0x2530((short) 0x2530, "Write error when accessing the DB"),
    EVENTID_0x2531((short) 0x2531, "Write error when accessing the DI"),
    EVENTID_0x2532((short) 0x2532, "Block number error when opening a DB"),
    EVENTID_0x2533((short) 0x2533, "Block number error when opening a DI"),
    EVENTID_0x2534((short) 0x2534, "Block number error when calling an FC"),
    EVENTID_0x2535((short) 0x2535, "Block number error when calling an FB"),
    EVENTID_0x253A((short) 0x253A, "DB not loaded"),
    EVENTID_0x253C((short) 0x253C, "FC not loaded"),
    EVENTID_0x253D((short) 0x253D, "SFC not loaded"),
    EVENTID_0x253E((short) 0x253E, "FB not loaded"),
    EVENTID_0x253F((short) 0x253F, "SFB not loaded"),
    EVENTID_0x2942((short) 0x2942, "I/O access error, reading"),
    EVENTID_0x2943((short) 0x2943, "I/O access error, writing"),

    
    //Event Class 3 - Asynchronous Errors
    EVENTID_0x3501((short) 0x3501, "Cycle time exceeded"),
    EVENTID_0x3502((short) 0x3502, "User interface (OB or FRB) request error"),
    EVENTID_0x3503((short) 0x3503, "Delay too long processing a priority class"),
    EVENTID_0x3505((short) 0x3505, "Time-of-day interrupt(s) skipped due to new clock setting"),
    EVENTID_0x3506((short) 0x3506, "Time-of-day interrupt(s) skipped when changing to RUN after HOLD"),
    EVENTID_0x3507((short) 0x3507, "Multiple OB request errors caused internal buffer overflow"),
    EVENTID_0x3508((short) 0x3508, "Synchronous cycle interrupt-timing error"),
    EVENTID_0x3509((short) 0x3509, "Interrupt loss due to excess interrupt load"),
    EVENTID_0x350A((short) 0x350A, "Resume RUN mode after CiR"),
    EVENTID_0x350B((short) 0x350B, "Technology synchronization interrupt - timing error"),
    EVENTID_0x3571((short) 0x3571, "Nesting depth too high in nesting levels"),
    EVENTID_0x3572((short) 0x3572, "Nesting depth for Master Control Relays too high"),
    EVENTID_0x3573((short) 0x3573, "Nesting depth too high after synchronous errors"),
    EVENTID_0x3574((short) 0x3574, "Nesting depth for block calls (U stack) too high"),
    EVENTID_0x3575((short) 0x3575, "Nesting depth for block calls (B stack) too high"),
    EVENTID_0x3576((short) 0x3576, "Local data allocation error"),
    EVENTID_0x3578((short) 0x3578, "Unknown instruction"),
    EVENTID_0x357A((short) 0x357A, "Jump instruction to target outside of the block"),
    EVENTID_0x3582((short) 0x3582, "Memory error detected and corrected by operating system"),
    EVENTID_0x3583((short) 0x3583, "Accumulation of detected and corrected memo errors"),
    EVENTID_0x3585((short) 0x3585, "Error in the PC operating system (only for LC RTX)"),
    EVENTID_0x3587((short) 0x3587, "Multi-bit memory error detected and corrected"),
    EVENTID_0x35A1((short) 0x35A1, "User interface (OB or FRB) not found"),
    EVENTID_0x35A2((short) 0x35A2, "OB not loaded (started by SFC or operating system due to configuration)"),
    EVENTID_0x35A3((short) 0x35A3, "Error when operating system accesses a block"),
    EVENTID_0x35A4((short) 0x35A4, "PROFInet Interface DB cannot be addressed"),
    EVENTID_0x35D2((short) 0x35D2, "Diagnostic entries cannot be sent at present"),
    EVENTID_0x35D3((short) 0x35D3, "Synchronization frames cannot be sent"),
    EVENTID_0x35D4((short) 0x35D4, "Illegal time jump resulting from synchronization"),
    EVENTID_0x35D5((short) 0x35D5, "Error adopting the synchronization time"),
    EVENTID_0x35E1((short) 0x35E1, "Incorrect frame ID in GD"),
    EVENTID_0x35E2((short) 0x35E2, "GD packet status cannot be entered in DB"),
    EVENTID_0x35E3((short) 0x35E3, "Frame length error in GD"),
    EVENTID_0x35E4((short) 0x35E4, "Illegal GD packet number received"),
    EVENTID_0x35E5((short) 0x35E5, "Error accessing DB in communication SFBs for configured S7 connections"),
    EVENTID_0x35E6((short) 0x35E6, "GD total status cannot be entered in DB"),
    EVENTID_0x3821((short) 0x3821, "BATTF: failure on at least one backup battery of the central rack, problem eliminated"),
    EVENTID_0x3822((short) 0x3822, "BAF: failure of backup voltage on central rack, problem eliminated"),
    EVENTID_0x3823((short) 0x3823, "24 volt supply failure on central rack, problem eliminated"),
    EVENTID_0x3825((short) 0x3825, "BATTF: failure on at least one backup battery of the redundant central rack, problem eliminated"),
    EVENTID_0x3826((short) 0x3826, "BAF: failure of backup voltage on redundant central rack, problem eliminated"),
    EVENTID_0x3827((short) 0x3827, "24 volt supply failure on redundant central rack, problem eliminated"),
    EVENTID_0x3831((short) 0x3831, "BATTF: failure of at least one backup battery of the expansion rack, problem eliminated"),
    EVENTID_0x3832((short) 0x3832, "BAF: failure of backup voltage on expansion rack, problem eliminated"),
    EVENTID_0x3833((short) 0x3833, "24 volt supply failure on at least one expansion rack, problem eliminated"),
    EVENTID_0x3842((short) 0x3842, "Module OK"),
    EVENTID_0x3854((short) 0x3854, "PROFINET IO interface submodule/submodule and matches the configured interface submodule/submodule"),
    EVENTID_0x3855((short) 0x3855, "PROFINET IO interface submodule/submodule inserted, but does not match the configured interface submodule/submodule"),
    EVENTID_0x3856((short) 0x3856, "PROFINET IO interface submodule/submodule inserted, but error in module parameter assignment"),
    EVENTID_0x3858((short) 0x3858, "PROFINET IO interface submodule access error corrected"),
    EVENTID_0x3861((short) 0x3861, "Module/interface module inserted, module type OK"),
    EVENTID_0x3863((short) 0x3863, "Module/interface module plugged in, but wrong module type"),
    EVENTID_0x3864((short) 0x3864, "Module/interface module plugged in, but causing problem (type ID unreadable)"),
    EVENTID_0x3865((short) 0x3865, "Module plugged in, but error in module parameter assignment"),
    EVENTID_0x3866((short) 0x3866, "Module can be addressed again, load voltage error removed"),
    EVENTID_0x3881((short) 0x3881, "Interface error leaving state"),
    EVENTID_0x3884((short) 0x3884, "Interface module plugged in"),
    EVENTID_0x38B3((short) 0x38B3, "I/O access error when updating the process image input table"),
    EVENTID_0x38B4((short) 0x38B4, "I/O access error when transferring the process image to the output modules"),
    EVENTID_0x38C1((short) 0x38C1, "Expansion rack operational again (1 to 21), leaving state"),
    EVENTID_0x38C2((short) 0x38C2, "Expansion rack operational again but mismatch between setpoint and actual configuration"),
    EVENTID_0x38C4((short) 0x38C4, "Distributed I/Os: station failure, leaving state"),
    EVENTID_0x38C5((short) 0x38C5, "Distributed I/Os: station fault, leaving state"),
    EVENTID_0x38C6((short) 0x38C6, "Expansion rack operational again, but error(s) in module parameter assignment"),
    EVENTID_0x38C7((short) 0x38C7, "DP: station operational again, but error(s) in module parameter assignment"),
    EVENTID_0x38C8((short) 0x38C8, "DP: station operational again, but mismatch between setpoint and actual configuration"),
    EVENTID_0x38CB((short) 0x38CB, "PROFINET IO station operational again"),
    EVENTID_0x38CC((short) 0x38CC, "PROFINET IO station error corrected"),
    EVENTID_0x3921((short) 0x3921, "BATTF: failure on at least one backup battery of the central rack"),
    EVENTID_0x3922((short) 0x3922, "BAF: failure of backup voltage on central rack"),
    EVENTID_0x3923((short) 0x3923, "24 volt supply failure on central rack"),
    EVENTID_0x3925((short) 0x3925, "BATTF: failure on at least one backup battery of the redundant central rack"),
    EVENTID_0x3926((short) 0x3926, "BAF: failure of backup voltage on redundant central rack"),
    EVENTID_0x3927((short) 0x3927, "24 volt supply failure on redundant central rack"),
    EVENTID_0x3931((short) 0x3931, "BATTF: failure of at least one backup battery of the expansion rack"),
    EVENTID_0x3932((short) 0x3932, "BAF: failure of backup voltage on expansion rack"),
    EVENTID_0x3933((short) 0x3933, "24 volt supply failure on at least one expansion rack"),
    EVENTID_0x3942((short) 0x3942, "Module error"),
    EVENTID_0x3951((short) 0x3951, "PROFINET IO submodule removed"),
    EVENTID_0x3954((short) 0x3954, "PROFINET IO interface submodule/submodule removed"),
    EVENTID_0x3961((short) 0x3961, "Module/interface module removed, cannot be addressed"),
    EVENTID_0x3966((short) 0x3966, "Module cannot be addressed, load voltage error"),
    EVENTID_0x3968((short) 0x3968, "Module reconfiguration has ended with error"),
    EVENTID_0x3981((short) 0x3981, "Interface error entering state"),
    EVENTID_0x3984((short) 0x3984, "Interface module removed"),
    EVENTID_0x3986((short) 0x3986, "Performance of an H-Sync link negatively affected"),
    EVENTID_0x39B1((short) 0x39B1, "I/O access error when updating the process image input table"),
    EVENTID_0x39B2((short) 0x39B2, "I/O access error when transferring the process image to the output modules"),
    EVENTID_0x39B3((short) 0x39B3, "I/O access error when updating the process image input table"),
    EVENTID_0x39B4((short) 0x39B4, "I/O access error when transferring the process image to the output modules"),
    EVENTID_0x39C1((short) 0x39C1, "Expansion rack failure (1 to 21), entering state"),
    EVENTID_0x39C3((short) 0x39C3, "Distributed I/Os: master system failure entering state"),
    EVENTID_0x39C4((short) 0x39C4, "Distributed I/Os: station failure, entering state"),
    EVENTID_0x39C5((short) 0x39C5, "Distributed I/Os: station fault, entering state"),
    EVENTID_0x39CA((short) 0x39CA, "PROFINET IO system failure"),
    EVENTID_0x39CB((short) 0x39CB, "PROFINET IO station failure"),
    EVENTID_0x39CC((short) 0x39CC, "PROFINET IO station error"),
    EVENTID_0x39CD((short) 0x39CD, "PROFINET IO station operational again, but expected configuration does not match actual configuration"),
    EVENTID_0x39CE((short) 0x39CE, "PROFINET IO station operational again, but error(s) in module parameter assignment"),
    EVENTID_0x42F3((short) 0x42F3, "Checksum error detected and corrected by the operating system"),

    
    //Event Class 4 - Stop Events and Other Mode Changes
    EVENTID_0x4300((short) 0x4300, "Backed-up power on"),
    EVENTID_0x4301((short) 0x4301, "Mode transition from STOP to STARTUP"),
    EVENTID_0x4302((short) 0x4302, "Mode transition from STARTUP to RUN"),
    EVENTID_0x4303((short) 0x4303, "STOP caused by stop switch being activated"),
    EVENTID_0x4304((short) 0x4304, "STOP caused by PG STOP operation or by SFB 20 STOP"),
    EVENTID_0x4305((short) 0x4305, "HOLD: breakpoint reached"),
    EVENTID_0x4306((short) 0x4306, "HOLD: breakpoint exited"),
    EVENTID_0x4307((short) 0x4307, "Memory reset started by PG operation"),
    EVENTID_0x4308((short) 0x4308, "Memory reset started by switch setting"),
    EVENTID_0x4309((short) 0x4309, "Memory reset started automatically (power on not backed up)"),
    EVENTID_0x430A((short) 0x430A, "HOLD exited, transition to STOP"),
    EVENTID_0x430D((short) 0x430D, "STOP caused by other CPU in multicomputing"),
    EVENTID_0x430E((short) 0x430E, "Memory reset executed"),
    EVENTID_0x430F((short) 0x430F, "STOP on the module due to STOP on a CPU"),
    EVENTID_0x4318((short) 0x4318, "Start of CiR"),
    EVENTID_0x4319((short) 0x4319, "CiR completed"),
    EVENTID_0x4357((short) 0x4357, "Module watchdog started"),
    EVENTID_0x4358((short) 0x4358, "All modules are ready for operation"),
    EVENTID_0x43B0((short) 0x43B0, "Firmware update was successful"),
    EVENTID_0x43B4((short) 0x43B4, "Error in firmware fuse"),
    EVENTID_0x43B6((short) 0x43B6, "Firmware updates canceled by redundant modules"),
    EVENTID_0x43D3((short) 0x43D3, "STOP on standby CPU"),
    EVENTID_0x43DC((short) 0x43DC, "Abort during link-up with switchover"),
    EVENTID_0x43DE((short) 0x43DE, "Updating aborted due to monitoring time being exceeded during the n-th attempt, new update attempt initiated"),
    EVENTID_0x43DF((short) 0x43DF, "Updating aborted for final time due to monitoring time being exceeded after completing the maximum amount of attempts. User intervention required"),
    EVENTID_0x43E0((short) 0x43E0, "Change from solo mode after link-up"),
    EVENTID_0x43E1((short) 0x43E1, "Change from link-up after updating"),
    EVENTID_0x43E2((short) 0x43E2, "Change from updating to redundant mode"),
    EVENTID_0x43E3((short) 0x43E3, "Master CPU: change from redundant mode to solo mode"),
    EVENTID_0x43E4((short) 0x43E4, "Standby CPU: change from redundant mode after error-search mode"),
    EVENTID_0x43E5((short) 0x43E5, "Standby CPU: change from error-search mode after link-up or STOP"),
    EVENTID_0x43E6((short) 0x43E6, "Link-up aborted on the standby CPU"),
    EVENTID_0x43E7((short) 0x43E7, "Updating aborted on the standby CPU"),
    EVENTID_0x43E8((short) 0x43E8, "Standby CPU: change from link-up after startup"),
    EVENTID_0x43E9((short) 0x43E9, "Standby CPU: change from startup after updating"),
    EVENTID_0x43F1((short) 0x43F1, "Reserve-master switchover"),
    EVENTID_0x43F2((short) 0x43F2, "Coupling of incompatible H-CPUs blocked by system program"),
    EVENTID_0x4510((short) 0x4510, "STOP violation of the CPU's data range"),
    EVENTID_0x4520((short) 0x4520, "DEFECTIVE: STOP not possible"),
    EVENTID_0x4521((short) 0x4521, "DEFECTIVE: failure of instruction processing processor"),
    EVENTID_0x4522((short) 0x4522, "DEFECTIVE: failure of clock chip"),
    EVENTID_0x4523((short) 0x4523, "DEFECTIVE: failure of clock pulse generator"),
    EVENTID_0x4524((short) 0x4524, "DEFECTIVE: failure of timer update function"),
    EVENTID_0x4525((short) 0x4525, "DEFECTIVE: failure of multicomputing synchronization"),
    EVENTID_0x4527((short) 0x4527, "DEFECTIVE: failure of I/O access monitoring"),
    EVENTID_0x4528((short) 0x4528, "DEFECTIVE: failure of scan time monitoring"),
    EVENTID_0x4530((short) 0x4530, "DEFECTIVE: memory test error in internal memory"),
    EVENTID_0x4532((short) 0x4532, "DEFECTIVE: failure of core resources"),
    EVENTID_0x4536((short) 0x4536, "DEFECTIVE: switch defective"),
    EVENTID_0x4540((short) 0x4540, "STOP: Memory expansion of the internal work memory has gaps. First memory expansion too small or missing"),
    EVENTID_0x4541((short) 0x4541, "STOP caused by priority class system"),
    EVENTID_0x4542((short) 0x4542, "STOP caused by object management system"),
    EVENTID_0x4543((short) 0x4543, "STOP caused by test functions"),
    EVENTID_0x4544((short) 0x4544, "STOP caused by diagnostic system"),
    EVENTID_0x4545((short) 0x4545, "STOP caused by communication system"),
    EVENTID_0x4546((short) 0x4546, "STOP caused by CPU memory management"),
    EVENTID_0x4547((short) 0x4547, "STOP caused by process image management"),
    EVENTID_0x4548((short) 0x4548, "STOP caused by I/O management"),
    EVENTID_0x454A((short) 0x454A, "STOP caused by configuration: an OB deselected with STEP 7 was being loaded into the CPU during STARTUP"),
    EVENTID_0x4550((short) 0x4550, "DEFECTIVE: internal system error"),
    EVENTID_0x4555((short) 0x4555, "No restart possible, monitoring time elapsed"),
    EVENTID_0x4556((short) 0x4556, "STOP: memory reset request from communication system / due to data inconsistency"),
    EVENTID_0x4562((short) 0x4562, "STOP caused by programming error (OB not loaded or not possible)"),
    EVENTID_0x4563((short) 0x4563, "STOP caused by I/O access error (OB not loaded or not possible)"),
    EVENTID_0x4567((short) 0x4567, "STOP caused by H event"),
    EVENTID_0x4568((short) 0x4568, "STOP caused by time error (OB not loaded or not possible)"),
    EVENTID_0x456A((short) 0x456A, "STOP caused by diagnostic interrupt (OB not loaded or not possible)"),
    EVENTID_0x456B((short) 0x456B, "STOP caused by removing/inserting module (OB not loaded or not possible)"),
    EVENTID_0x456C((short) 0x456C, "STOP caused by CPU hardware error (OB not loaded or not possible, or no FRB)"),
    EVENTID_0x456D((short) 0x456D, "STOP caused by program sequence error (OB not loaded or not possible)"),
    EVENTID_0x456E((short) 0x456E, "STOP caused by communication error (OB not loaded or not possible)"),
    EVENTID_0x456F((short) 0x456F, "STOP caused by rack failure OB (OB not loaded or not possible)"),
    EVENTID_0x4570((short) 0x4570, "STOP caused by process interrupt (OB not loaded or not possible)"),
    EVENTID_0x4571((short) 0x4571, "STOP caused by nesting stack error"),
    EVENTID_0x4572((short) 0x4572, "STOP caused by master control relay stack error"),
    EVENTID_0x4573((short) 0x4573, "STOP caused by exceeding the nesting depth for synchronous errors"),
    EVENTID_0x4574((short) 0x4574, "STOP caused by exceeding interrupt stack nesting depth in the priority class stack"),
    EVENTID_0x4575((short) 0x4575, "STOP caused by exceeding block stack nesting depth in the priority class stack"),
    EVENTID_0x4576((short) 0x4576, "STOP caused by error when allocating the local data"),
    EVENTID_0x4578((short) 0x4578, "STOP caused by unknown opcode"),
    EVENTID_0x457A((short) 0x457A, "STOP caused by code length error"),
    EVENTID_0x457B((short) 0x457B, "STOP caused by DB not being loaded on on-board I/Os"),
    EVENTID_0x457D((short) 0x457D, "Reset/clear request because the version of the internal interface to the integrated technology was changed"),
    EVENTID_0x457F((short) 0x457F, "STOP caused by STOP command"),
    EVENTID_0x4580((short) 0x4580, "STOP: back-up buffer contents inconsistent (no transition to RUN)"),
    EVENTID_0x4590((short) 0x4590, "STOP caused by overloading the internal functions"),
    EVENTID_0x45D5((short) 0x45D5, "LINK-UP rejected due to mismatched CPU memory configuration of the sub-PLC"),
    EVENTID_0x45D6((short) 0x45D6, "LINK-UP rejected due to mismatched system program of the sub-PLC"),
    EVENTID_0x45D8((short) 0x45D8, "DEFECTIVE: hardware fault detected due to other error"),
    EVENTID_0x45D9((short) 0x45D9, "STOP due to SYNC module error"),
    EVENTID_0x45DA((short) 0x45DA, "STOP due to synchronization error between H CPUs"),
    EVENTID_0x45DD((short) 0x45DD, "LINK-UP rejected due to running test or other online functions"),
    EVENTID_0x4926((short) 0x4926, "DEFECTIVE: failure of the watchdog for I/O access"),
    EVENTID_0x4931((short) 0x4931, "STOP or DEFECTIVE: memory test error in memory submodule"),
    EVENTID_0x4933((short) 0x4933, "Checksum error"),
    EVENTID_0x4934((short) 0x4934, "DEFECTIVE: memory not available"),
    EVENTID_0x4935((short) 0x4935, "DEFECTIVE: cancelled by watchdog/processor exceptions"),
    EVENTID_0x4949((short) 0x4949, "STOP caused by continuous hardware interrupt"),
    EVENTID_0x494D((short) 0x494D, "STOP caused by I/O error"),
    EVENTID_0x494E((short) 0x494E, "STOP caused by power failure"),
    EVENTID_0x494F((short) 0x494F, "STOP caused by configuration error"),
    EVENTID_0x4959((short) 0x4959, "One or more modules not ready for operation"),
    EVENTID_0x497C((short) 0x497C, "STOP caused by integrated technology"),
    EVENTID_0x49A0((short) 0x49A0, "STOP caused by parameter assignment error or non-permissible variation of setpoint and actual extension: Start-up blocked"),
    EVENTID_0x49A1((short) 0x49A1, "STOP caused by parameter assignment error: memory reset request"),
    EVENTID_0x49A2((short) 0x49A2, "STOP caused by error in parameter modification: startup disabled"),
    EVENTID_0x49A3((short) 0x49A3, "STOP caused by error in parameter modification: memory reset request"),
    EVENTID_0x49A4((short) 0x49A4, "STOP: inconsistency in configuration data"),
    EVENTID_0x49A5((short) 0x49A5, "STOP: distributed I/Os: inconsistency in the loaded configuration information"),
    EVENTID_0x49A6((short) 0x49A6, "STOP: distributed I/Os: invalid configuration information"),
    EVENTID_0x49A7((short) 0x49A7, "STOP: distributed I/Os: no configuration information"),
    EVENTID_0x49A8((short) 0x49A8, "STOP: error indicated by the interface module for the distributed I/Os"),
    EVENTID_0x49B1((short) 0x49B1, "Firmware update data incorrect"),
    EVENTID_0x49B2((short) 0x49B2, "Firmware update: hardware version does not match firmware"),
    EVENTID_0x49B3((short) 0x49B3, "Firmware update: module type does not match firmware"),
    EVENTID_0x49D0((short) 0x49D0, "LINK-UP aborted due to violation of coordination rules"),
    EVENTID_0x49D1((short) 0x49D1, "LINK-UP/UPDATE sequence aborted"),
    EVENTID_0x49D2((short) 0x49D2, "Standby CPU changed to STOP due to STOP on the master CPU during link-up"),
    EVENTID_0x49D4((short) 0x49D4, "STOP on a master, since partner CPU is also a master (link-up error)"),
    EVENTID_0x49D7((short) 0x49D7, "LINK-UP rejected due to change in user program or in configuration"),
    EVENTID_0x42F4((short) 0x42F4, "Standby CPU: connection/update via SFC90 is locked in the master CPU"),
            
    //Event Class 5 - Mode Run-time Events
    EVENTID_0x530D((short) 0x530D, "New startup information in the STOP mode"),
    EVENTID_0x510F((short) 0x510F, "A problem as occurred with WinLC. This problem has caused the CPU to go into STOP mode or has caused a fault in the CPU"),    
    EVENTID_0x5311((short) 0x5311, "Startup despite Not Ready message from module(s)"),
    EVENTID_0x5371((short) 0x5371, "Distributed I/Os: end of the synchronization with a DP master"),
    EVENTID_0x5380((short) 0x5380, "Diagnostic buffer entries of interrupt and asynchronous errors disabled"),
    EVENTID_0x5395((short) 0x5395, "Distributed I/Os: reset of a DP master"),
    EVENTID_0x53A2((short) 0x53A2, "Download of technology firmware successful"),
    EVENTID_0x53A4((short) 0x53A4, "Download of technology DB not successful"),
    EVENTID_0x5445((short) 0x5445, "Start of System reconfiguration in RUN mode"),
    EVENTID_0x5481((short) 0x5481, "All licenses for runtime software are complete again"),
    EVENTID_0x5498((short) 0x5498, "No more inconsistency with DP master systems due to CiR"),
    EVENTID_0x5545((short) 0x5545, "Start of System reconfiguration in RUN mode"),
    EVENTID_0x5581((short) 0x5581, "One or several licenses for runtime software are missing"),
    EVENTID_0x558A((short) 0x558A, "Difference between the MLFB of the configured and inserted CPU"),
    EVENTID_0x558B((short) 0x558B, "Difference in the firmware version of the configured and inserted CPU"),
    EVENTID_0x5598((short) 0x5598, "Start of possible inconsistency with DP master systems due to CiR"),
    EVENTID_0x55A5((short) 0x55A5, "Version conflict: internal interface with integrated technology"),
    EVENTID_0x55A6((short) 0x55A6, "The maximum number of technology objects has been exceeded"),
    EVENTID_0x55A7((short) 0x55A7, "A technology DB of this type is already present"),
    EVENTID_0x5879((short) 0x5879, "Diagnostic message from DP interface: EXTF LED off"),
    EVENTID_0x5960((short) 0x5960, "Parameter assignment error when switching"),
    EVENTID_0x5961((short) 0x5961, "Parameter assignment error"),
    EVENTID_0x5962((short) 0x5962, "Parameter assignment error preventing startup"),
    EVENTID_0x5963((short) 0x5963, "Parameter assignment error with memory reset request"),
    EVENTID_0x5966((short) 0x5966, "Parameter assignment error when switching"),
    EVENTID_0x5969((short) 0x5969, "Parameter assignment error with startup blocked"),
    EVENTID_0x596A((short) 0x596A, "PROFINET IO: IP address of an IO device already present"),
    EVENTID_0x596B((short) 0x596B, "IP address of an Ethernet interface already exists"),
    EVENTID_0x596C((short) 0x596C, "Name of an Ethernet interface already exists"),
    EVENTID_0x596D((short) 0x596D, "The existing network configuration does not mach the system requirements or configuration"),
    EVENTID_0x5979((short) 0x5979, "Diagnostic message from DP interface: EXTF LED on"),
    EVENTID_0x597C((short) 0x597C, "DP Global Control command failed or moved"),
    EVENTID_0x59A0((short) 0x59A0, "The interrupt can not be associated in the CPU"),
    EVENTID_0x59A1((short) 0x59A1, "Configuration error in the integrated technology"),
    EVENTID_0x59A3((short) 0x59A3, "Error when downloading the integrated technology"),
    EVENTID_0x53FF((short) 0x53FF, "Reset to factory setting"),
    

    
    //Event Class 6 - Communication Events
    EVENTID_0x6316((short) 0x6316, "Interface error when starting programmable controller"),
    EVENTID_0x6353((short) 0x6353, "Firmware update: Start of firmware download over the network"),
    EVENTID_0x6390((short) 0x6390, "Formatting of Micro Memory Card complete"),
    EVENTID_0x6500((short) 0x6500, "Connection ID exists twice on module"),
    EVENTID_0x6501((short) 0x6501, "Connection resources inadequate"),
    EVENTID_0x6502((short) 0x6502, "Error in the connection description"),
    EVENTID_0x6510((short) 0x6510, "CFB structure error detected in instance DB when evaluating EPROM"),
    EVENTID_0x6514((short) 0x6514, "GD packet number exists twice on the module"),
    EVENTID_0x6515((short) 0x6515, "Inconsistent length specifications in GD configuration information"),
    EVENTID_0x6521((short) 0x6521, "No memory submodule and no internal memory available"),
    EVENTID_0x6522((short) 0x6522, "Illegal memory submodule: replace submodule and reset memory"),
    EVENTID_0x6523((short) 0x6523, "Memory reset request due to error accessing submodule"),
    EVENTID_0x6524((short) 0x6524, "Memory reset request due to error in block header"),
    EVENTID_0x6526((short) 0x6526, "Memory reset request due to memory replacement"),
    EVENTID_0x6527((short) 0x6527, "Memory replaced, therefore restart not possible"),
    EVENTID_0x6528((short) 0x6528, "Object handling function in the STOP/HOLD mode, no restart possible"),
    EVENTID_0x6529((short) 0x6529, "No startup possible during the load"),
    EVENTID_0x652A((short) 0x652A, "No startup because block exists twice in user memory user program function"),
    EVENTID_0x652B((short) 0x652B, "No startup because block is too long for submodule - replace submodule"),
    EVENTID_0x652C((short) 0x652C, "No startup due to illegal OB on submodule"),
    EVENTID_0x6532((short) 0x6532, "No startup because illegal configuration information on submodule"),
    EVENTID_0x6533((short) 0x6533, "Memory reset request because of invalid submodule content"),
    EVENTID_0x6534((short) 0x6534, "No startup: block exists more than once on submodule"),
    EVENTID_0x6535((short) 0x6535, "No startup: not enough memory to transfer block from submodule"),
    EVENTID_0x6536((short) 0x6536, "No startup: submodule contains an illegal block number"),
    EVENTID_0x6537((short) 0x6537, "No startup: submodule contains a block with an illegal length"),
    EVENTID_0x6538((short) 0x6538, "Local data or write-protection ID (for DB) of a block illegal for CPU"),
    EVENTID_0x6539((short) 0x6539, "Illegal command in block (detected by compiler)"),
    EVENTID_0x653A((short) 0x653A, "Memory reset request because local OB data on submodule too short"),
    EVENTID_0x6543((short) 0x6543, "No startup: illegal block type"),
    EVENTID_0x6544((short) 0x6544, "No startup: attribute relevant for processing illegal"),
    EVENTID_0x6545((short) 0x6545, "Source language illegal"),
    EVENTID_0x6546((short) 0x6546, "Maximum amount of configuration information reached"),
    EVENTID_0x6547((short) 0x6547, "Parameter assignment error assigning parameters to modules (not on P bus, cancel download)"),
    EVENTID_0x6548((short) 0x6548, "Plausibility error during block check"),
    EVENTID_0x6549((short) 0x6549, "Structure error in block"),
    EVENTID_0x6550((short) 0x6550, "A block has an error in the CRC"),
    EVENTID_0x6551((short) 0x6551, "A block has no CRC"),
    EVENTID_0x6253((short) 0x6253, "Firmware update: End of firmware download over the network"),    
    EVENTID_0x6560((short) 0x6560, "SCAN overflow"),
    EVENTID_0x6805((short) 0x6805, "Resource problem on configured connections, eliminated"),
    EVENTID_0x6881((short) 0x6881, "Interface error leaving state"),
    EVENTID_0x6905((short) 0x6905, "Resource problem on configured connections"),
    EVENTID_0x6981((short) 0x6981, "Interface error entering state"),
    
    //Event Class 7 - H/F Events
    EVENTID_0x72A2((short) 0x72A2, "Failure of a DP master or a DP master system"),
    EVENTID_0x72A3((short) 0x72A3, "Redundancy restored on the DP slave"),
    EVENTID_0x72DB((short) 0x72DB, "Safety program: safety mode disabled"),
    EVENTID_0x72E0((short) 0x72E0, "Loss of redundancy in communication, problem eliminated"),
    EVENTID_0x7301((short) 0x7301, "Loss of redundancy (1 of 2) due to failure of a CPU"),
    EVENTID_0x7302((short) 0x7302, "Loss of redundancy (1 of 2) due to STOP on the standby triggered by user"),
    EVENTID_0x7303((short) 0x7303, "H system (1 of 2) changed to redundant mode"),
    EVENTID_0x7323((short) 0x7323, "Discrepancy found in operating system data"),
    EVENTID_0x7331((short) 0x7331, "Standby-master switchover due to master failure"),
    EVENTID_0x7333((short) 0x7333, "Standby-master switchover due to system modification during runtime"),
    EVENTID_0x7334((short) 0x7334, "Standby-master switchover due to communication error at the synchronization module"),
    EVENTID_0x7340((short) 0x7340, "Synchronization error in user program due to elapsed wait time"),
    EVENTID_0x7341((short) 0x7341, "Synchronization error in user program due to waiting at different synchronization points"),
    EVENTID_0x7342((short) 0x7342, "Synchronization error in operating system due to waiting at different synchronization points"),
    EVENTID_0x7343((short) 0x7343, "Synchronization error in operating system due to elapsed wait time"),
    EVENTID_0x7344((short) 0x7344, "Synchronization error in operating system due to incorrect data"),
    EVENTID_0x734A((short) 0x734A, "The \"Re-enable\" job triggered by SFC 90 \"H_CTRL\" was executed"),
    EVENTID_0x73A3((short) 0x73A3, "Loss of redundancy on the DP slave"),
    EVENTID_0x73C1((short) 0x73C1, "Update process canceled"),
    EVENTID_0x73C2((short) 0x73C2, "Updating aborted due to monitoring time being exceeded during the n-th attempt (1 = n = max. possible number of update attempts after abort due to excessive monitoring time)"),
    EVENTID_0x73D8((short) 0x73D8, "Safety mode disabled"),
    EVENTID_0x73DB((short) 0x73DB, "Safety program: safety mode enabled"),
    EVENTID_0x73E0((short) 0x73E0, "Loss of redundancy in communication"),
    EVENTID_0x74DD((short) 0x74DD, "Safety program: Shutdown of a fail-save runtime group disabled"),
    EVENTID_0x74DE((short) 0x74DE, "Safety program: Shutdown of the F program disabled"),
    EVENTID_0x74DF((short) 0x74DF, "Start of F program initialization"),
    EVENTID_0x7520((short) 0x7520, "Error in RAM comparison"),
    EVENTID_0x7521((short) 0x7521, "Error in comparison of process image output value"),
    EVENTID_0x7522((short) 0x7522, "Error in comparison of memory bits, timers, or counters"),
    EVENTID_0x75D1((short) 0x75D1, "Safety program: Internal CPU error"),
    EVENTID_0x75D2((short) 0x75D2, "Safety program error: Cycle time time-out"),
    EVENTID_0x75D6((short) 0x75D6, "Data corrupted in safety program prior to the output to F I/O"),
    EVENTID_0x75D7((short) 0x75D7, "Data corrupted in safety program prior to the output to partner F-CPU"),
    EVENTID_0x75D9((short) 0x75D9, "Invalid REAL number in a DB"),
    EVENTID_0x75DA((short) 0x75DA, "Safety program: Error in safety data format"),
    EVENTID_0x75DC((short) 0x75DC, "Runtime group, internal protocol error"),
    EVENTID_0x75DD((short) 0x75DD, "Safety program: Shutdown of a fail-save runtime group enabled"),
    EVENTID_0x75DE((short) 0x75DE, "Safety program: Shutdown of the F program enabled"),
    EVENTID_0x75DF((short) 0x75DF, "End of F program initialization"),
    EVENTID_0x75E1((short) 0x75E1, "Safety program: Error in FB \"F_PLK\" or \"F_PLK_O\" or \"F_CYC_CO\" or \"F_TEST\" or \"F_TESTC\""),
    EVENTID_0x75E2((short) 0x75E2, "Safety program: Area length error"),
    EVENTID_0x7852((short) 0x7852, "SYNC module inserted"),
    EVENTID_0x7855((short) 0x7855, "SYNC module eliminated"),
    EVENTID_0x78D3((short) 0x78D3, "Communication error between PROFIsafe and F I/O"),
    EVENTID_0x78D4((short) 0x78D4, "Error in safety relevant communication between F CPUs"),
    EVENTID_0x78D5((short) 0x78D5, "Error in safety relevant communication between F CPUs"),
    EVENTID_0x78E3((short) 0x78E3, "F-I/O device input channel depassivated"),
    EVENTID_0x78E4((short) 0x78E4, "F-I/O device output channel depassivated"),
    EVENTID_0x78E5((short) 0x78E5, "F-I/O device depassivated"),
    EVENTID_0x7934((short) 0x7934, "Standby-master switchover due to connection problem at the SYNC module"),
    EVENTID_0x7950((short) 0x7950, "Synchronization module missing"),
    EVENTID_0x7951((short) 0x7951, "Change at the SYNC module without Power On"),
    EVENTID_0x7952((short) 0x7952, "SYNC module removed"),
    EVENTID_0x7953((short) 0x7953, "Change at the SYNC-module without reset"),
    EVENTID_0x7954((short) 0x7954, "SYNC module: rack number assigned twice"),
    EVENTID_0x7955((short) 0x7955, "SYNC module error"),
    EVENTID_0x7956((short) 0x7956, "Illegal rack number set on SYNC module"),
    EVENTID_0x7960((short) 0x7960, "Redundant I/O: Time-out of discrepancy time at digital input, error is not yet localized"),
    EVENTID_0x7961((short) 0x7961, "Redundant I/O, digital input error: Signal change after expiration of the discrepancy time"),
    EVENTID_0x7962((short) 0x7962, "Redundant I/O: Digital input error"),
    EVENTID_0x796F((short) 0x796F, "Redundant I/O: The I/O was globally disabled"),
    EVENTID_0x7970((short) 0x7970, "Redundant I/O: Digital output error"),
    EVENTID_0x7980((short) 0x7980, "Redundant I/O: Time-out of discrepancy time at analog input"),
    EVENTID_0x7981((short) 0x7981, "Redundant I/O: Analog input error"),
    EVENTID_0x7990((short) 0x7990, "Redundant I/O: Analog output error"),
    EVENTID_0x79D3((short) 0x79D3, "Communication error between PROFIsafe and F I/O"),
    EVENTID_0x79D4((short) 0x79D4, "Error in safety relevant communication between F CPUs"),
    EVENTID_0x79D5((short) 0x79D5, "Error in safety relevant communication between F CPUs"),
    EVENTID_0x79E3((short) 0x79E3, "F-I/O device input channel passivated"),
    EVENTID_0x79E4((short) 0x79E4, "F-I/O device output channel passivated"),
    EVENTID_0x79E5((short) 0x79E5, "F-I/O device passivated"),
    EVENTID_0x79E6((short) 0x79E6, "Inconsistent safety program"),
    EVENTID_0x79E7((short) 0x79E7, "Simulation block (F system block) loaded"),
    EVENTID_0x73E8((short) 0x73E8, "Consistency of the safety program verified by testing"),
    EVENTID_0x73E9((short) 0x73E9, "Consistency of the safety program cannot be checked"),
    
    //Event Class 8 - Diagnostic Events for Modules
    EVENTID_0x8x00((short) 0x8000, "Module fault/OK"),
    EVENTID_0x8x01((short) 0x8001, "Internal error"),
    EVENTID_0x8x02((short) 0x8002, "External error"),
    EVENTID_0x8x03((short) 0x8003, "Channel error"),
    EVENTID_0x8x04((short) 0x8004, "External error"),
    EVENTID_0x8x05((short) 0x8005, "No front connector"),
    EVENTID_0x8x06((short) 0x8006, "No parameter assignment"),
    EVENTID_0x8x07((short) 0x8007, "Incorrect parameters in module"),  
    EVENTID_0x8x30((short) 0x8030, "User submodule incorrect/not found"),
    EVENTID_0x8x31((short) 0x8031, "Communication problem"),
    EVENTID_0x8x32((short) 0x8032, "Operating mode: RUN/STOP (STOP: entering state, RUN: leaving state)"),
    EVENTID_0x8x33((short) 0x8033, "Time monitoring responded (watchdog)"),
    EVENTID_0x8x34((short) 0x8034, "Internal module power failure"),
    EVENTID_0x8x35((short) 0x8035, "BATTF: battery exhausted"),
    EVENTID_0x8x36((short) 0x8036, "Total backup failed"),
    EVENTID_0x8x40((short) 0x8040, "Expansion rack failed"),
    EVENTID_0x8x41((short) 0x8041, "Processor failure"),
    EVENTID_0x8x42((short) 0x8042, "EPROM error"),
    EVENTID_0x8x43((short) 0x8043, "RAM error"),
    EVENTID_0x8x44((short) 0x8044, "ADC/DAC error"),
    EVENTID_0x8x45((short) 0x8045, "Fuse blown"),
    EVENTID_0x8x46((short) 0x8046, "Hardware interrupt lost"),    
    EVENTID_0x8x50((short) 0x8050, "Configuration/parameter assignment error. Analog input"), 
    EVENTID_0x8x51((short) 0x8051, "Common mode error"), 
    EVENTID_0x8x52((short) 0x8052, "Short circuit to phase"),
    EVENTID_0x8x53((short) 0x8053, "Short circuit to ground"),
    EVENTID_0x8x54((short) 0x8054, "Wire break"),
    EVENTID_0x8x55((short) 0x8055, "Reference channel error"),
    EVENTID_0x8x56((short) 0x8056, "Below measuring range"),
    EVENTID_0x8x57((short) 0x8057, "Above measuring range"),
    EVENTID_0x8x60((short) 0x8060, "Configuration/parameter assignment error. Analog output"),
    EVENTID_0x8x61((short) 0x8061, "Common mode error"),
    EVENTID_0x8x62((short) 0x8062, "Short circuit to phase"),
    EVENTID_0x8x63((short) 0x8063, "Short circuit to ground"),
    EVENTID_0x8x64((short) 0x8064, "Wire break"),
    EVENTID_0x8x66((short) 0x8066, "No load voltage"),
    EVENTID_0x8x70((short) 0x8070, "Configuration/parameter assignment error. Digital input"),
    EVENTID_0x8x71((short) 0x8071, "Chassis ground fault"),
    EVENTID_0x8x72((short) 0x8072, "Short circuit to phase (sensor)"),
    EVENTID_0x8x73((short) 0x8073, "Short circuit to ground (sensor)"),
    EVENTID_0x8x74((short) 0x8074, "Wire break"),
    EVENTID_0x8x75((short) 0x8075, "No sensor power supply"),
    EVENTID_0x8x80((short) 0x8080, "Configuration/parameter assignment error. Digital output"),
    EVENTID_0x8x81((short) 0x8081, "Chassis ground fault"),
    EVENTID_0x8x82((short) 0x8082, "Short circuit to phase"),
    EVENTID_0x8x83((short) 0x8083, "Short circuit to ground"),
    EVENTID_0x8x84((short) 0x8084, "Wire break"),
    EVENTID_0x8x85((short) 0x8085, "Fuse tripped"),
    EVENTID_0x8x86((short) 0x8086, "No load voltage"),
    EVENTID_0x8x87((short) 0x8087, "Excess temperature"),
    EVENTID_0x8xB0((short) 0x80B0, "Counter module, signal A faulty. FM"),
    EVENTID_0x8xB1((short) 0x80B1, "Counter module, signal B faulty"),
    EVENTID_0x8xB2((short) 0x80B2, "Counter module, signal N faulty"),
    EVENTID_0x8xB3((short) 0x80B3, "Counter module, incorrect value passed between the channels"),
    EVENTID_0x8xB4((short) 0x80B4, "Counter module, 5.2 V sensor supply faulty"),
    EVENTID_0x8xB5((short) 0x80B5, "Counter module, 24 V sensor supply faulty"),
    
    //Event Class 9 - Standard User Events
    EVENTID_0x9001((short) 0x9001, "Automatic mode"),
    EVENTID_0x9101((short) 0x9101, "Manual mode"),
    EVENTID_0x9x02((short) 0x9002, "OPEN/CLOSED, ON/OFF"),
    EVENTID_0x9x03((short) 0x9003, "Manual command enable"),
    EVENTID_0x9x04((short) 0x9004, "Unit protective command (OPEN/CLOSED)"),
    EVENTID_0x9x05((short) 0x9005, "Process enable"),
    EVENTID_0x9x06((short) 0x9006, "System protection command"),
    EVENTID_0x9x07((short) 0x9007, "Process value monitoring responded"),
    EVENTID_0x9x08((short) 0x9008, "Manipulated variable monitoring responded"),
    EVENTID_0x9x09((short) 0x9009, "System deviation greater than permitted"),
    EVENTID_0x9x0A((short) 0x900A, "Limit position error"),
    EVENTID_0x9x0B((short) 0x900B, "Runtime error"),
    EVENTID_0x9x0C((short) 0x900C, "Command execution error (sequencer)"),
    EVENTID_0x9x0D((short) 0x900D, "Operating status running > OPEN"),
    EVENTID_0x9x0E((short) 0x900E, "Operating status running > CLOSED"),
    EVENTID_0x9x0F((short) 0x900F, "Command blocking"),
    EVENTID_0x9x11((short) 0x9011, "Process status OPEN/ON"),
    EVENTID_0x9x12((short) 0x9012, "Process status CLOSED/OFF"),
    EVENTID_0x9x13((short) 0x9013, "Process status intermediate position"),
    EVENTID_0x9x14((short) 0x9014, "Process status ON via AUTO"),
    EVENTID_0x9x15((short) 0x9015, "Process status ON via manual"),
    EVENTID_0x9x16((short) 0x9016, "Process status ON via protective command"),
    EVENTID_0x9x17((short) 0x9017, "Process status OFF via AUTO"),
    EVENTID_0x9x18((short) 0x9018, "Process status OFF via manual"),
    EVENTID_0x9x19((short) 0x9019, "Process status OFF via protective command"),
    EVENTID_0x9x21((short) 0x9021, "Function error on approach"),
    EVENTID_0x9x22((short) 0x9022, "Function error on leaving"),
    EVENTID_0x9x31((short) 0x9031, "Actuator (DE/WE) limit position OPEN"),
    EVENTID_0x9x32((short) 0x9032, "Actuator (DE/WE) limit position not OPEN"),
    EVENTID_0x9x33((short) 0x9033, "Actuator (DE/WE) limit position CLOSED"),
    EVENTID_0x9x34((short) 0x9034, "Actuator (DE/WE) limit position not CLOSED"),
    EVENTID_0x9x41((short) 0x9041, "Illegal status, tolerance time elapsed"),
    EVENTID_0x9x42((short) 0x9042, "Illegal status, tolerance time not elapsed"),
    EVENTID_0x9x43((short) 0x9043, "Interlock error, tolerance time = 0"),
    EVENTID_0x9x44((short) 0x9044, "Interlock error, tolerance time > 0"),
    EVENTID_0x9x45((short) 0x9045, "No reaction"),
    EVENTID_0x9x46((short) 0x9046, "Final status exited illegally, tolerance time = 0"),
    EVENTID_0x9x47((short) 0x9047, "Final status exited illegally, tolerance time > 0"),
    EVENTID_0x9x50((short) 0x9050, "Upper limit of signal range USR"),
    EVENTID_0x9x51((short) 0x9051, "Upper limit of measuring range UMR"),
    EVENTID_0x9x52((short) 0x9052, "Lower limit of signal range LSR"),
    EVENTID_0x9x53((short) 0x9053, "Lower limit of measuring range LMR"),
    EVENTID_0x9x54((short) 0x9054, "Upper alarm limit UAL"),
    EVENTID_0x9x55((short) 0x9055, "Upper warning limit UWL"),
    EVENTID_0x9x56((short) 0x9056, "Upper tolerance limit UTL"),
    EVENTID_0x9x57((short) 0x9057, "Lower tolerance limit LTL"),
    EVENTID_0x9x58((short) 0x9058, "Lower warning limit LWL"),
    EVENTID_0x9x59((short) 0x9059, "Lower alarm limit LAL"),
    EVENTID_0x9x60((short) 0x9060, "GRAPH7 step entering/leaving"),
    EVENTID_0x9x61((short) 0x9061, "GRAPH7 interlock error"),
    EVENTID_0x9x62((short) 0x9062, "GRAPH7 execution error"),
    EVENTID_0x9x63((short) 0x9063, "GRAPH7 error noted"),
    EVENTID_0x9x64((short) 0x9064, "GRAPH7 error acknowledged"),
    EVENTID_0x9x70((short) 0x9070, "Trend exceeded in positive direction"),
    EVENTID_0x9x71((short) 0x9071, "Trend exceeded in negative direction"),
    EVENTID_0x9x72((short) 0x9072, "No reaction"),
    EVENTID_0x9x73((short) 0x9073, "Final state exited illegally"),
    EVENTID_0x9x80((short) 0x9080, "Limit value exceeded, tolerance time = 0"),
    EVENTID_0x9x81((short) 0x9081, "Limit value exceeded, tolerance time > 0"),
    EVENTID_0x9x82((short) 0x9082, "Below limit value, tolerance time = 0"),
    EVENTID_0x9x83((short) 0x9083, "Below limit value, tolerance time > 0"),
    EVENTID_0x9x84((short) 0x9084, "Gradient exceeded, tolerance time = 0"),
    EVENTID_0x9x85((short) 0x9085, "Gradient exceeded, tolerance time > 0"),
    EVENTID_0x9x86((short) 0x9086, "Below gradient, tolerance time = 0"),
    EVENTID_0x9x87((short) 0x9087, "Below gradient, tolerance time > 0"),
    EVENTID_0x9090((short) 0x9090, "User parameter assignment error entering/leaving"),
    EVENTID_0x9190((short) 0x9190, "User parameter assignment error entering/leaving"),
    EVENTID_0x91F0((short) 0x91F0, "Overflow"),
    EVENTID_0x91F1((short) 0x91F1, "Underflow"),
    EVENTID_0x91F2((short) 0x91F2, "Division by 0"),
    EVENTID_0x91F3((short) 0x91F3, "Illegal calculation operation"),
    
    //Event Classes A and B - Free User Events
    
    
    //Event Classes C,D & E - Reserved Event Classes
    
    //Event Classes F -  Reserved for modules not in central 
    //                   rack (for example, CPs or FMs)
    
    EVENTID_0x0000((short) 0x0000, "NULL: Check for user information.");

    
    
    private static final Map<Short, S7DiagnosticEventId> map;
    
    private static final Map<Short, String> idstr;
    
    static {
        map = new HashMap<>();
        idstr = new HashMap<>();
        for (S7DiagnosticEventId  event : S7DiagnosticEventId.values()) {
            map.put(event.code, event);
        }
       
        idstr.put((short) 0x0000, "Event leaving state. ");
        idstr.put((short) 0x0100, "Event entering state. ");
        idstr.put((short) 0x0200, "Entry in diagnostic buffer. ");
        idstr.put((short) 0x0400, "Internal error. ");
        idstr.put((short) 0x0800, "External error. ");
    }    
    
    private final String description;
    private final short code;
    
    S7DiagnosticEventId(final short code, final String description){
        this.code = code;        
        this.description = description;
    }
    
    public String getDescription(){
        //short id = (short) (code & 0x0F00);
        return description;
    }    
    
    public short getCode() {
        return code;
    }    

    public static S7DiagnosticEventId  valueOf(short code) {
        
        Integer intcode = Short.toUnsignedInt(code);
        
        if ((intcode > 0x8000) && (intcode < 0xA000)){
            if ((intcode != 0x9101) && (intcode != 0x9190) && (intcode != 0x91F0)
                    && (intcode != 0x91F1) && (intcode != 0x91F2)
                    && (intcode != 0x91F3))
            intcode = (intcode & 0xF0FF);
        } else if (((intcode >= 0xA000) && (intcode <= 0xBFFF))) intcode = 0x0000;
                
        return map.get(intcode.shortValue());
    }    
}
