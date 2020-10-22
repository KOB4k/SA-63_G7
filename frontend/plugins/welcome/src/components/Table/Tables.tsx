import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';
import MenuIcon from '@material-ui/icons/Menu';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Link from '@material-ui/core/Link';
import {
    Content,
    ContentHeader,
} from '@backstage/core';
import { Link as RouterLink } from 'react-router-dom';
import {
    Button,
    AppBar,
    Toolbar,
    Typography,
    IconButton
} from '@material-ui/core';


const useStyles = makeStyles(theme => ({
    table: {
        minWidth: 650,
    },
    title: {
        flexGrow: 1,
    },
    menuButton: {
        marginRight: theme.spacing(2),
    },
    logoutButton: {
        marginLeft: 10,
        marginRight: 10,
        color: 'white'
    }
}));

function redirecLogOut() {
    // redire Page ... http://localhost:3000/
    window.location.href = "http://localhost:3000/";
}
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [diseases, setDisease] = useState(Array);
 const [loading, setLoading] = useState(true);

 useEffect(() => {
   const getDiseases = async () => {
     const res = await api.listDisease({ limit: 10, offset: 0 });
     setLoading(false);
     setDisease(res);
   };
   getDiseases();
 }, [loading]);
 

 
 const deleteDiseases = async (id: number) => {
   const res = await api.deleteDisease({ id: id });
   setLoading(true);
 };
 
 return (
   
  <div>
  <AppBar position="static">
      <Toolbar>
          <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
              <MenuIcon />
          </IconButton>
          <Typography variant="h4" className={classes.title}>
              ระบบจัดการโรคติดต่อ
</Typography>
          <IconButton
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              color="inherit"
              size="medium"
          >
              <AccountCircle />
              <Typography>
                  <Link variant="h6" onClick={redirecLogOut} className={classes.logoutButton}>
                      LOGOUT
  </Link>
              </Typography>
          </IconButton>
      </Toolbar>
  </AppBar>
  <Content>
  <ContentHeader title="">
      <Button
          size="large"
          style={{ float: 'right', marginBottom: 'auto' }}
          color="primary"
          component={RouterLink}
          to="/Disease"
          variant="contained"
      >
          เพิ่มข้อมูลโรคติดต่อ
   </Button>
   </ContentHeader>

   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">Employee</TableCell>
           <TableCell align="center">Name</TableCell>
           <TableCell align="center">Severity</TableCell>
           <TableCell align="center">Symptom</TableCell>
           <TableCell align="center">Contagion</TableCell>
           <TableCell align="center">Diseasetype</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {diseases.map((item:any) => (
           <TableRow key={item.id}>
             
             <TableCell align="center">{item.edges?.employee?.userId}</TableCell>
             <TableCell align="center">{item.name}</TableCell>
             <TableCell align="center">{item.edges?.severity?.name}</TableCell>
             <TableCell align="center">{item.symptom}</TableCell>
             <TableCell align="center">{item.contagion}</TableCell>
             <TableCell align="center">{item.edges?.diseasetype?.name}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteDiseases(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
</div>
 );
}
