import './app.scss'
import {RouterProvider} from "react-router-dom"
import {router} from "@/router/router"

export default function App() {

    return (
        <>
            {/*<Notifications/>*/}
            <RouterProvider router={router}/>
        </>

    )
}
