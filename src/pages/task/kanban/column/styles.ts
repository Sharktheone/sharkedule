import {createStyles, rem} from "@mantine/core"


export const useStyles = createStyles((theme) => ({
    column: {
        backgroundColor: theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.white,
        padding: "1rem",
        borderRadius: theme.radius.md,
        border: `1px solid ${
            theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2]
        }`,
        width: "24rem",
        boxShadow: theme.shadows.md,
    },
    title: {
        marginBottom: rem(20),
    },
    add: {
        "& > div > textarea": {
            backgroundColor: theme.colorScheme === "dark" ? theme.colors.dark[7] : theme.colors.gray[0],
            borderColor: theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2],
            shadow: theme.shadows.md,
        },
        shadow: theme.shadows.md,
    },
    ghost: {
        border: `1px solid ${
            theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2]
        }`,
    }
}))