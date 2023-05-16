import {createStyles, rem} from "@mantine/core";


export const useStyles = createStyles((theme) => ({
    column: {
        backgroundColor: theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.white,
        padding: "1rem",
        borderRadius: theme.radius.md,
        border: `1px solid ${
            theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2]
        }`,
        width: "24rem",
        shadow: theme.shadows.md,
    },
    title: {
        marginBottom: rem(20),
    }
}))