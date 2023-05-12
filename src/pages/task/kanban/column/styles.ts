import {createStyles, rem} from "@mantine/core";


export const useStyles = createStyles((theme) => ({
    column: {
        backgroundColor: theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.white,
        padding: "1rem",
        borderRadius: theme.radius.sm,
        border: `1px solid ${
            theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2]
        }`,
        width: "24rem",
    },
    title: {
        borderRadius: theme.radius.sm,
        marginBottom: rem(20),

    }
}))