import {createStyles} from "@mantine/core"


export const useColors = createStyles((theme) => ({
    menu: {
        backgroundColor: theme.colors.dark[6],
        border: `1px solid ${theme.colors.dark[5]}`
    }

}))