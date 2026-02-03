import { createRouter , createWebHistory} from 'vue-router';
import UserPage from '../pages/UserPage.vue';
import LandingPage from '../pages/LandingPage.vue';
import StringerPage from '../pages/StringerPage.vue'
import AdminPage from '../pages/AdminPage.vue'

const routes = [
  {
    path: '/',
    name: 'LandingPage',
    component: LandingPage
  },

  {
    path: '/UserPage',
    name: 'UserPage',
    component: UserPage
  }, 

  {
    path: '/StringerPage',
    name: 'StringerPage',
    component: StringerPage
  },

  {
    path: '/AdminPage',
    name: 'AdminPage',
    component: AdminPage
  }
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
});

export default router;