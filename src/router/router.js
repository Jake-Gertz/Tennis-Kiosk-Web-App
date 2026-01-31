import { createRouter , createWebHistory} from 'vue-router';
import UserPage from '../pages/UserPage.vue';
import LandingPage from '../pages/LandingPage.vue';

const routes = [
  {
    path: '/UserPage',
    name: 'UserPage',
    component: UserPage
  }, 

  {
    path: '/',
    name: 'LandingPage',
    component: LandingPage
  }
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
});

export default router;