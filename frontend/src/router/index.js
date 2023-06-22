// Composables
import { createRouter, createWebHistory } from 'vue-router'

const routes = [{
    path: "/",
    component: () =>
        import ("@/layouts/default/Default.vue"),
    children: [{
            path: "",
            name: "Products",
            component: () =>
                import ("@/views/Products.vue"),
        },
        {
            path: "/success",
            name: "Success",
            component: () =>
                import ("@/views/Success.vue"),
        },
        {
            path: "/canceled",
            name: "Cancel",
            component: () =>
                import ("@/views/Cancel.vue"),
        },
    ],
}, ];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
})

export default router