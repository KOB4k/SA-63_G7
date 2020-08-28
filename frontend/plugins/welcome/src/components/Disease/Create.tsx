import React, { useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import {TextField,InputAdornment} from '@material-ui/core';
import Autocomplete from '@material-ui/lab/Autocomplete';
import AccountCircle from '@material-ui/icons/AccountCircle';

const useStyles = makeStyles((theme: Theme) =>
 createStyles({
   root: {
     display: 'flex',
     flexWrap: 'wrap',
     justifyContent: 'center',
   },
   margin: {
     margin: theme.spacing(3),
   },
   withoutLabel: {
     marginTop: theme.spacing(3),
   },
   textField: {
     width: '25ch',
   },
 }),
);

const diseaseState = {
  emp_id: 1,
  name: '',
  severity_id: 1,
  manner: '',
  contagion: '',
  agency_id: [1,2] 
 };

const employeeState = [
  {
    emp_id: 1,
    name: 'toei'
  },
  {
    emp_id: 2,
    name: 'kob'
  }
];
const angencyState = [
  {
    angency_id: 1,
    name: 'มทส'
  },
  {
    angency_id: 2,
    name: 'มหาราช'
  }
];
const severityState = [
  {
    severity_id: 1,
    name: 'เริ่มต้น'
  },
  {
    severity_id: 2,
    name: 'รุนแรง'
  },
  {
    severity_id: 3,
    name: 'รุนแรงมาก'
  }
];
export default function Create() {
 const classes = useStyles();
 const profile = { givenName: 'to Software Analysis 63' };
 const api = new DefaultApi();
 
 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);
 const [disease, setDisease] = useState(diseaseState);
 const [employee, setEmployee] = useState(employeeState);
 const [angency, setAngency] = useState(angencyState);
 const [severity, setSeverity] = useState(severityState);

 const handleInputDiseaseChange = (event: any) => {
   const { id, value } = event.target;
   setDisease({ ...disease, [id]: value });
 };
 
 const CreateUser = async () => {
  //  const res:any = await api.createUser({ user });
  //  setStatus(true);
  //  if (res.id != ''){
  //    setAlert(true);
  //  } else {
  //    setAlert(false);
  //  }
  // const timer = setTimeout(() => {
  //   setStatus(false);
  // }, 1000);
  setStatus(true);
  setAlert(false);
   const timer = setTimeout(() => {
     setStatus(false);
   }, 1000);
 };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`ระบบจัดการโรคติดต่อ`}
       subtitle="Some quick intro and links."
     ></Header>
     <Content>
       <ContentHeader title="เพิ่มข้อมูลโรคติดต่อ">
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 This is a success alert — check it out!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 This is a warning alert — check it out!
               </Alert>
             )}
           </div>
         ) : null}
       </ContentHeader>
       <div className={classes.root}>
         <form noValidate autoComplete="off">
           <FormControl
             fullWidth
             className={classes.margin}
             variant="outlined"
           >
            <Autocomplete
              id="combo-box-demo"
              options={employee}
              getOptionLabel={(option) =>option.emp_id + ' : ' + option.name}
              renderInput={(params) => <TextField {...params} label="รหัสพนักงาน" variant="outlined" />}
            />
             <TextField
                style={{ marginTop: 20 }}
               id="name"
               label="ชื่อโรค"
               variant="outlined"
               type="string"
               size="medium"
               value={disease.name}
               onChange={handleInputDiseaseChange}
             />
            <Autocomplete
                style={{ marginTop: 20 }}
              id="combo-box-demo"
              options={angency}
              getOptionLabel={(option) => option.angency_id + ' - ' + option.name}
              renderInput={(params) => <TextField {...params} label="ระดับความรุนแรง" variant="outlined" />}
            />

           </FormControl>
           

           <div className={classes.margin}>
             <Button
               onClick={() => {
                 CreateUser();
               }}
               variant="contained"
               color="primary"
             >
               Submit
             </Button>
             <Button
               style={{ marginLeft: 20 }}
               component={RouterLink}
               to="/"
               variant="contained"
             >
               Back
             </Button>
           </div>
         </form>
       </div>
     </Content>
   </Page>
 );
}
