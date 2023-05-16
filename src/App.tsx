import './App.scss'
import {RouterProvider} from "react-router-dom"
import {router} from "./router/router"
import {MantineProvider} from "@mantine/core"
import {Notifications} from "@mantine/notifications"

function App() {

    return (
        <MantineProvider withGlobalStyles withNormalizeCSS theme={{
            colorScheme: "dark",
        }}>
            <Notifications/>
            <RouterProvider router={router}/>
        </MantineProvider>

    )
}

export default App
