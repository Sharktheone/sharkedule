import './App.scss'
import {RouterProvider} from "react-router-dom"
import {router} from "./router/router"
import {MantineProvider} from "@mantine/core"

function App() {

    return (
        <MantineProvider withGlobalStyles withNormalizeCSS theme={{
            colorScheme: "dark",
        }}>
            <RouterProvider router={router}/>
        </MantineProvider>

    )
}

export default App
