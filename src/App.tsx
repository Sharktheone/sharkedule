import '@/App.scss'
import {RouterProvider} from "react-router-dom"
import {router} from "@/router/router"
import {createTheme, MantineProvider} from "@mantine/core"
import {Notifications} from "@mantine/notifications"


const theme = createTheme({
    fontFamily: 'sans-serif',
    primaryColor: 'orange',

})

function App() {

    return (
        <MantineProvider theme={theme}>
            <Notifications/>
            <RouterProvider router={router}/>
        </MantineProvider>

    )
}

export default App
