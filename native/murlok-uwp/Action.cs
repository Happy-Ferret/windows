using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Windows.Foundation.Collections;

namespace murlok_uwp
{
    enum ActionType
    {
        DriverLaunched,
        DriverTerminating,
        DriverTerminated
    }

    class Action
    {
        public static ValueSet New(ActionType type)
        {
            return New(type, "");
        }

        public static ValueSet New(ActionType type, string payload)
        {
            ValueSet valueSet = new ValueSet();
            valueSet.Add("type", type.ToString());
            valueSet.Add("payload", payload);
            return valueSet;
        }
    }
}
