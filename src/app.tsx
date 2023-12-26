import "./app.scss"
import "styles/color-modifiers.scss"
import {RouterProvider} from "react-router-dom"
import {router} from "@/router/router"
import {ToastContainer} from "react-toastify"

export default function App() {

    return (
        <>
            <ToastContainer position="bottom-left" autoClose={5000} newestOnTop draggable theme="dark" pauseOnHover/>
            <RouterProvider router={router}/>
        </>

    )
}
