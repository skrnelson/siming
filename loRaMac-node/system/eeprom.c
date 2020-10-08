/*!
 * \file      eeprom.c
 *
 * \brief     EEPROM driver implementation
 *
 * \copyright Revised BSD License, see section \ref LICENSE.
 *
 * \code
 *                ______                              _
 *               / _____)             _              | |
 *              ( (____  _____ ____ _| |_ _____  ____| |__
 *               \____ \| ___ |    (_   _) ___ |/ ___)  _ \
 *               _____) ) ____| | | || |_| ____( (___| | | |
 *              (______/|_____)_|_|_| \__)_____)\____)_| |_|
 *              (C)2013-2017 Semtech
 *
 * \endcode
 *
 * \author    Miguel Luis ( Semtech )
 *
 * \author    Gregory Cristian ( Semtech )
 */
#include "eeprom.h"

static uint8_t EepromMcuWriteBuffer( uint16_t addr, uint8_t *buffer, uint16_t size )
{
    return 0;
}

static uint8_t EepromMcuReadBuffer( uint16_t addr, uint8_t *buffer, uint16_t size )
{
    return 0;
}

uint8_t EepromWriteBuffer( uint16_t addr, uint8_t *buffer, uint16_t size )
{
    // Boundary check
    if( size > ( 0xFFFF - addr ) )
    {
        return 0;
    }
    return EepromMcuWriteBuffer( addr, buffer, size );
}

uint8_t EepromReadBuffer( uint16_t addr, uint8_t *buffer, uint16_t size )
{
    // Boundary check
    if( size > ( 0xFFFF - addr ) )
    {
        return 0;
    }
    return EepromMcuReadBuffer( addr, buffer, size );
}

void EepromSetDeviceAddr( uint8_t addr )
{
    // EepromMcuSetDeviceAddr( addr );
}

uint8_t EepromGetDeviceAddr( void )
{
    // return EepromMcuGetDeviceAddr( );
    return 0;
}