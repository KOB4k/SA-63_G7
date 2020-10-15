import { createPlugin } from '@backstage/core';
import Disease from './components/Disease'


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Disease);

  },
});
