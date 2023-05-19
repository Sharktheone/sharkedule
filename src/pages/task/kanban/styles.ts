import {createStyles} from "@mantine/core"


export const useStyles = createStyles((theme) => ({
    addColumn: {
        borderRadius: theme.radius.md,
        backgroundColor: theme.colorScheme === "dark" ? theme.colors.dark[4] : theme.white,
        border: `1px solid ${
            theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2]
        }`,
        shadow: theme.shadows.md,
    }
}))