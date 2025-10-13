interface WizardPageProps { 
    onSelect?: (info: any) => void; 
    onLogin?: (info: any) => void;
    onBack?: () => void; 
    disk?: any;
    partition?: any;
}

export type { WizardPageProps };