 
import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreateDisease from './components/Disease';
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/disease', CreateDisease);
  },
});
 
