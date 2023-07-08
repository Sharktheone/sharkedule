import {createStyles} from "@mantine/core"


export const useColors = createStyles((theme) => ({
    description: {
        backgroundColor: theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.colors.gray[7],
        border: `1px solid ${
            theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2]
        }`,
        "--outline-color": theme.colors.blue[5],
    },

}))