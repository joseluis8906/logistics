import React, { useState } from 'react'
import Container from '@mui/material/Container'
import { Card, CardContent } from '@mui/material'
import TextField from '@mui/material/TextField'
import Box from '@mui/material/Box'
import { TableContainer, Table, TableHead, TableBody, TableRow, TableCell } from '@mui/material'
import Paper from '@mui/material/Paper'
import Typography from '@mui/material/Typography'

export default function Home () {
  const guardMinVal = (evt: React.ChangeEvent<HTMLInputElement>, cb: (a: number) => void) => {
    const current: number = Number(evt.target.value)
    if (current >= 0) {
      cb(current)
    }
  }

  const [width, setWidth] = useState(0)
  const controlWidth = (evt: React.ChangeEvent<HTMLInputElement>) => guardMinVal(evt, setWidth)

  const [height, setHeight] = useState(0)
  const controlHeight = (evt: React.ChangeEvent<HTMLInputElement>) => guardMinVal(evt, setHeight)

  const [length, setLength] = useState(0)
  const controlLength = (evt: React.ChangeEvent<HTMLInputElement>) => guardMinVal(evt, setLength)

  const [weight, setWeight] = useState(0)
  const controlWeight = (evt: React.ChangeEvent<HTMLInputElement>) => guardMinVal(evt, setWeight)

  return (
    <Container sx={{ mt: 6 }}>
      <Card sx={{ mx: 'auto', p: 2, width: 500, height: 'auto' }}>
        <CardContent><Typography>Calculate Shipping</Typography></CardContent>
        <CardContent>
          <Box component='form' sx={{ '& > :not(style)': { m: 1, width: '27ch' } }} noValidate autoComplete='off'>
            <TextField label="From" type="text"/>
            <TextField label="To" type="text"/>
            <TextField label="Width" type="text" onChange={ controlWidth } value={ width }/>
            <TextField label="Height" type="text" onChange={ controlHeight } value={ height }/>
            <TextField label="Length" type="text" onChange={ controlLength } value={ length }/>
            <TextField label="Weight" type="text" onChange={ controlWeight } value={ weight }/>
          </Box>
          <TableContainer component={ Paper }>
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell>Description</TableCell>
                  <TableCell align='right'>Amount</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                <TableRow>
                  <TableCell>Distance</TableCell>
                  <TableCell align='right'>$0</TableCell>
                </TableRow>
                <TableRow>
                  <TableCell>Size</TableCell>
                  <TableCell align='right'>$0</TableCell>
                </TableRow>
                <TableRow>
                  <TableCell>Weight</TableCell>
                  <TableCell align='right'>$0</TableCell>
                </TableRow>
                <TableRow>
                  <TableCell>Total</TableCell>
                  <TableCell align='right'>$0</TableCell>
                </TableRow>
              </TableBody>
            </Table>
          </TableContainer>
        </CardContent>
      </Card>
    </Container>
  )
}
