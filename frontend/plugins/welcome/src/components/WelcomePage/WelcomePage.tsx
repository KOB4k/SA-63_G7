import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 const student = { name: '',lastname: 'poonkasem' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`ระบบจัดการโรคติดต่อ`}
       subtitle="Some quick intro and links."
     ></Header>
     <Content>
       <ContentHeader title="ระบบเก็บข้อมูลโรคติดต่อ">
         <Link component={RouterLink} to="/disease">
           <Button variant="contained" color="primary">
             เพิ่มข้อมูล
           </Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;
 
