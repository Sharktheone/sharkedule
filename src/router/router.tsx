import {createBrowserRouter} from "react-router-dom"
import KanbanBoardLoader from "@/pages/task/kanban/loader"
import boardsLoader from "@/pages/task/boardsloader"
import {lazy} from "react"
import Home from "@/pages/home/home"
import Login from "@/pages/login/login"


const KanbanBoards = lazy(() => import("@/pages/task/kanbanboards"))
const BoardsError = lazy(() => import("@/pages/task/boardserror"))
const Kanban = lazy(() => import("@/pages/task/kanban/kanban"))
export const router = createBrowserRouter([
    {
        path: "/",
        element: <Home/>,
    },
    {
        path: "/dashboard",
    },
    {
        path: "/login",
        element: <Login/>,
    },
    {
        path: "/register",
    },
    {
        path: "/board",
        element: <KanbanBoards/>,
        loader: boardsLoader,
        errorElement: <BoardsError/>,
        children: [
            {
                path: "assigned"
            },
            {
                path: "new"
            },
        ]
    },
    {
        path: "/board/:uuid",
        loader: KanbanBoardLoader,
        element: <Kanban/>,
    },
    {
        path: "/calendar",
        children: [
            {
                path: "new"
            },
        ]
    },
    {
        path: "/tickets",
        children: [
            {
                path: "assigned"
            },
            {
                path: "new"
            },
            {
                path: ":uuid"
            }

        ]
    },
    {
        path: "/admin",
        children: [
            {
                path: "users"
            },
            {
                path: "roles"
            },
            {
                path: "permissions"
            },
            {
                path: "settings"
            }
        ]
    }
])