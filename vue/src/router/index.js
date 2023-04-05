import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '@/components/Login'
import Home from '@/components/Home'
import Welcome from '@/components/Welcome'
import Reserve from '@/components/Reserve'
import AdminApply from '@/components/AdminApply'
import MeetingRoom from '@/components/MeetingRoom'
import AppointmentRecord from '@/components/AppointmentRecord'
import Opinion from '@/components/Opinion'
import RoomAnalysis from '@/components/RoomAnalysis'
import PeriodAnalysis from "@/components/PeriodAnalysis";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/CanteenPC/login'
  },
  { path: '/CanteenPC/login',
    component: Login
  },
  {
    path: '/CanteenPC/home',
    component: Home,
    redirect: '/CanteenPC/welcome',
    children: [
      { path: '/CanteenPC/Welcome', component: Welcome },
      { path: '/CanteenPC/Reserve', component: Reserve },
      { path: '/CanteenPC/AdminApply', component: AdminApply },
      { path: '/CanteenPC/MeetingRoom', component: MeetingRoom },
      { path: '/CanteenPC/AppointmentRecord', component: AppointmentRecord },
      { path: '/CanteenPC/Opinion', component: Opinion },
      { path: '/CanteenPC/RoomAnalysis', component: RoomAnalysis },
      { path: '/CanteenPC/PeriodAnalysis', component: PeriodAnalysis }

    ]
  }

]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})


export default router
