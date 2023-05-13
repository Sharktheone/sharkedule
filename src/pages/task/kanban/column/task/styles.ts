import {createStyles} from "@mantine/core"


export const useStyles = createStyles((theme) => ({
    task: {
        minHeight: "3rem",
        backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white,
        border: `1px solid ${
            theme.colorScheme === 'dark' ? theme.colors.dark[5] : theme.colors.gray[2]
        }`,
        borderRadius: theme.radius.sm,
        padding: theme.spacing.xs,
        marginBottom: theme.spacing.xs,
        boxShadow: theme.shadows.sm,
    }
}))