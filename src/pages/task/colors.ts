import {createStyles} from "@mantine/core"


export const useColors = createStyles((theme) => ({
    colors: {
        backgroundColor: theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.white,

        "& > ul > li:not(.no-boards)": {
            backgroundColor: theme.colorScheme === "dark" ? theme.colors.dark[7] : theme.white,
            border: `1px solid ${theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2]}`,
            "& > a": {
                color: theme.colorScheme === "dark" ? theme.colors.dark[0] : theme.colors.gray[9],
            },
            "&:hover": {
                backgroundColor: theme.colorScheme === "dark" ? theme.colors.dark[8] : theme.colors.gray[0],
            }
        },

    }
}))