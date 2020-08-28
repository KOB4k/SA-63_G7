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
import { TextField, InputAdornment } from '@material-ui/core';
import Autocomplete from '@material-ui/lab/Autocomplete';
import DeleteIcon from '@material-ui/icons/Delete';
import SaveIcon from '@material-ui/icons/Save';
import ArrowBackIosIcon from '@material-ui/icons/ArrowBackIos';
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
      width: '50ch',
    },
    button: {
      marginLeft:'40px',
    },
    center: {
      textAlign: 'center',
    }
  }),
);

const diseaseState = {
  emp_id: 1,
  name: '',
  severity_id: 1,
  symptom: '',
  contagion: '',
  agency_id: [1, 2]
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


const typeState = [
  {
    severity_id: 1,
    name: 'โรคติดต่อ'
  },
  {
    severity_id: 2,
    name: 'โรคติดต่อต้องแจ้งความ'
  },
  {
    severity_id: 3,
    name: 'โรคติดต่ออันตราย'
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
  const [severity, setSeverity] = useState(severityState);
  const [type, setType] = useState(typeState);

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
        <ContentHeader title=''> 
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
        <h2 className={classes.center}>เพิ่มข้อมูลโรคติดต่อ</h2>
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
                getOptionLabel={(option) => option.emp_id + ' : ' + option.name}
                renderInput={(params) => <TextField {...params} label="รหัสพนักงาน" variant="outlined" />}
              />

              <TextField
                style={{ marginTop: 20 }}
                id="name"
                label="ชื่อโรค"
                variant="outlined"
                type="string"
                size="medium"
                multiline
                value={disease.name}
                onChange={handleInputDiseaseChange}
              />

              <Autocomplete
                style={{ marginTop: 20 }}
                id="combo-box-demo"
                options={severity}
                getOptionLabel={(option) => option.severity_id + ' - ' + option.name}
                renderInput={(params) => <TextField {...params} label="ระดับความรุนแรง" variant="outlined" />}
              />

              <TextField
                style={{ marginTop: 20 }}
                id="symptom"
                label="อาการ"
                variant="outlined"
                type="string"
                size="medium"
                multiline
                rows={4}
                value={disease.symptom}
                onChange={handleInputDiseaseChange}
              />

              <TextField
                style={{ marginTop: 20 }}
                id="contagion"
                label="การแพร่กระจาย"
                variant="outlined"
                multiline
                rows={4}
                type="string"
                size="medium"
                value={disease.contagion}
                onChange={handleInputDiseaseChange}
              />

              <Autocomplete
                style={{ marginTop: 20 }}
                id="combo-box-demo"
                options={type}
                getOptionLabel={(option) => option.name}
                renderInput={(params) => <TextField {...params} label="ประเภทโรคติดต่อ" variant="outlined" />}
              />

            </FormControl>
            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreateUser();
                }}
                variant="contained"
                color="primary"
                size="small"
                className={classes.button}
                startIcon={<SaveIcon />}
              >
                Save
              </Button>
              <Button
                variant="contained"
                color="secondary"
                size="small"
                className={classes.button}
                startIcon={<DeleteIcon />}
              >
                RESET
              </Button>
              <Button
                variant="contained"
                color="blueGrey"
                size="small"
                component={RouterLink}
                to="/"
                className={classes.button}
                startIcon={<ArrowBackIosIcon />}
              >
                BACK
              </Button>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}
