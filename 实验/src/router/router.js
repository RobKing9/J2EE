import Vue from "vue"
import Router from "vue-router"

import HelloWorld from '@/components/HelloWorld'
import Home from '@/components/Home'
import One from '@/components/one'
import Two from '@/components/two'
import Three from '@/components/three'


Vue.use(Router)

export default new Router({
	mode:'history',
	routes:[
		{
			path:'/',
			name:'HelloWorld',
			component:HelloWorld
		},
		{
			path:'/Home',
			name:'Home',
			component:Home,
			children:[
				{
					path:'one',
					component:One
				},
				{
					path:'two',
					component:Two
				},
				{
					path:'three',
					component:Three
				}
			]
		}	
	]
})

