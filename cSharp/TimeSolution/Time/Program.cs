using System;

namespace Time
{
    class Program
    {
        static void Main(string[] args)
        {

            //To access arguments while compiling in VS17
            //In Solution Explorer, right click project, click 'Properties'
            //Click 'Debug'
            //enter arguments, and Ctrl-S to save
            //Return to program, in Debug Menu, select "Start without debugging"
            if(args.Length > 0)
            {
                Console.WriteLine("Hello, " + args[0]);
            }
            else
            {
                Console.WriteLine("Hello World!");
            }

            Console.WriteLine("How many hours of sleep did you get last night?");
            int hoursOfSleep = int.Parse(Console.ReadLine());
            
            //DateTime package under the System Library
            Console.WriteLine("Today is: " + DateTime.Now.DayOfWeek);
            Console.WriteLine("And the time is: " + DateTime.Now.TimeOfDay);

            if (hoursOfSleep > 8)
            {
                Console.WriteLine("You should be well rested");
            }
            else
            {
                Console.WriteLine("You need more sleep");
            }

            Console.ReadLine();
        }
    }
}
