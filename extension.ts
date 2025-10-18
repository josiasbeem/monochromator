import * as vscode from 'vscode';

const THEME_MIGRATION_MAP: Record<string, string> = {
    'Monochromator Dark+': 'Monochromator Dark',
    'Monochromator Dark+ Amber': 'Monochromator Dark Amber',
    'Monochromator Dark+ Amethyst': 'Monochromator Dark Amethyst',
    'Monochromator Dark+ Aquamarine': 'Monochromator Dark Aquamarine',
    'Monochromator Dark+ Emerald': 'Monochromator Dark Emerald',
    'Monochromator Dark+ Plain': 'Monochromator Dark Plain',
    'Monochromator Dark+ Ruby': 'Monochromator Dark Ruby',
    'Monochromator Dark+ Sulfur': 'Monochromator Dark Sulfur'
};

function migrateTheme() {
    const config = vscode.workspace.getConfiguration();
    let migrated = false;

    // Migrate workbench.colorTheme
    const currentTheme = config.get<string>('workbench.colorTheme');
    if (currentTheme && THEME_MIGRATION_MAP[currentTheme]) {
        const newTheme = THEME_MIGRATION_MAP[currentTheme];
        config.update('workbench.colorTheme', newTheme, vscode.ConfigurationTarget.Global);
        migrated = true;
    }

    // Migrate workbench.preferredDarkColorTheme
    const preferredDark = config.get<string>('workbench.preferredDarkColorTheme');
    if (preferredDark && THEME_MIGRATION_MAP[preferredDark]) {
        const newTheme = THEME_MIGRATION_MAP[preferredDark];
        config.update('workbench.preferredDarkColorTheme', newTheme, vscode.ConfigurationTarget.Global);
        migrated = true;
    }

    if (migrated) {
        vscode.window.showInformationMessage('Monochromator: Dark+ theme migrated to Dark');
    }
}

export function activate(context: vscode.ExtensionContext) {
    migrateTheme();
}
