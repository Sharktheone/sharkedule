import {createStyles} from "@mantine/core"


export const useColors = createStyles((theme) => ({
    colors: {
        backgroundColor: theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.white,

        "& > ul > li > a": {
            backgroundColor: theme.colorScheme === "dark" ? theme.colors.dark[7] : theme.white,
            borderColor: theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2],
            color: theme.colorScheme === "dark" ? theme.colors.dark[0] : theme.colors.gray[9],
        },

    }
}))