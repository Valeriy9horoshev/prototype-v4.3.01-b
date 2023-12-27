import java.util.Map;
import java.util.Map.Entry;
import org.wesnoth.utils.ResourceUtils;
import org.wesnoth.utils.WorkspaceUtils;

    private static final String PREPROCESSED_FILE_PATH = WorkspaceUtils
                                                           .getTemporaryFolder( )
                                                           + "preprocessed.txt";


    public static PreprocessorUtils getInstance( )
    {
        return PreprocessorUtilsInstance.instance_;
    }

        if( ResourcesPlugin.getWorkspace( ).validateLinkLocation( coreLibrary,
            paths.getCoreDirPath( ) ).getCode( ) != IStatus.ERROR ) {
            try {

                    }

                    coreLibrary.delete( true, new NullProgressMonitor( ) );
                }

    public static IProject createWesnothProject( String name, String location,
        String installName, IProgressMonitor monitor )
    {
        return createWesnothProject( name, location, installName, false,
            monitor );
    }
