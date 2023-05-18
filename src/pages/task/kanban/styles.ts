import {createStyles} from "@mantine/core"


export const useStyles = createStyles((theme) => ({
    addColumn: {
        backgroundColor: theme.colorScheme === "dark" ? theme.colors.dark[4] : theme.white,
        borderRadius: theme.radius.md,
        border: `1px solid ${
            theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2]
        }`,
        width: "24rem",
        shadow: theme.shadows.md,
    }
}))