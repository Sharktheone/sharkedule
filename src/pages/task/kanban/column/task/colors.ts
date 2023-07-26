import {createStyles} from "@mantine/core"

export const useColors = createStyles((theme) => ({
    task: {
        backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white,
        "&:hover": {
            backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[8] : theme.colors.gray[0],
        },
        border: `1px solid ${
            theme.colorScheme === 'dark' ? theme.colors.dark[5] : theme.colors.gray[2]
        }`,
        boxShadow: theme.shadows.sm,
    }
}))