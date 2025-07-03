package com.miamorito.android.ui

import androidx.compose.foundation.layout.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.unit.dp
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.compose.rememberNavController
import com.miamorito.android.ui.screens.CharacterListScreen
import com.miamorito.android.ui.screens.ChatScreen
import com.miamorito.android.utils.DeviceUtils

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun MiAmoritoApp() {
    val navController = rememberNavController()
    val context = LocalContext.current
    val deviceId = remember { DeviceUtils.getDeviceId(context) }

    NavHost(
        navController = navController,
        startDestination = "character_list"
    ) {
        composable("character_list") {
            CharacterListScreen(
                deviceId = deviceId,
                onCharacterSelected = { characterId ->
                    navController.navigate("chat/$characterId")
                }
            )
        }
        
        composable("chat/{characterId}") { backStackEntry ->
            val characterId = backStackEntry.arguments?.getString("characterId") ?: ""
            ChatScreen(
                deviceId = deviceId,
                characterId = characterId,
                onBackClick = {
                    navController.popBackStack()
                }
            )
        }
    }
}