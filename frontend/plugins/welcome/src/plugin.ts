import { createPlugin } from '@backstage/core';
import Disease from './components/Disease'
import Login from './components/Login'


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/Disease', Disease);

  },
});
