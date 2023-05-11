import {DataRouteObject, RouteObject} from "react-router-dom"


export type route = {
    permission?: string,
    authentication_needed: boolean,
    children?: route[],

} &Omit<RouteObject, 'children'>


export let routes: route[] = [
    {
        path: "/",
        authentication_needed: false,
    },
    {
        path: "/login",
        authentication_needed: false,
    },
    {
        path: "/register",
        authentication_needed: false,
    },
    {
        path: "/dashboard",
        authentication_needed: true,
    },
    {
        path: "/task",
        authentication_needed: true,


        children: [
            {
                path: "assigned",
                authentication_needed: true,
            },
            {
                path: "new",
                authentication_needed: true,
            },
            {
                path: ":uuid",
                authentication_needed: true,
            }
        ]
    },
    {
        path: "/calendar",
        authentication_needed: true,
        children: [
            {
                path: "new",
                authentication_needed: true,
            },
        ]
    },
    {
        path: "/tickets",
        authentication_needed: true,
        children: [
            {
                path: "assigned",
                authentication_needed: true,
            },
            {
                path: "new",
                authentication_needed: false,
            },
            {
                path: ":uuid",
                authentication_needed: true,
            }

        ]
    },
    {
        path: "/admin",
        authentication_needed: true,
        children: [
            {
                path: "users",
                authentication_needed: true,
            },
            {
                path: "roles",
                authentication_needed: true,
            },
            {
                path: "permissions",
                authentication_needed: true,
            },
            {
                path: "settings",
                authentication_needed: true,
            }
        ]
    }
]