// Composables
import { createRouter, createWebHistory } from 'vue-router'

const routes = [{
        path: "/",
        name: "Products",
        component: () =>
            import ("@/views/Products.vue"),
        meta: {
            layout: "Default",
        },
    },
    {
        path: "/success",
        name: "Success",
        component: () =>
            import ("@/views/Success.vue"),
        meta: {
            layout: "Default",
        },
    },
    {
        path: "/failed",
        name: "Failure",
        component: () =>
            import ("@/views/Failure.vue"),
        meta: {
            layout: "Default",
        },
    },
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
})

export default router